package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/yakuzzaa/timeTracker/internal/api/repository"
	"github.com/yakuzzaa/timeTracker/internal/api/serializers"
	"github.com/yakuzzaa/timeTracker/internal/models"
)

type UserService struct {
	repo   repository.User
	logger *slog.Logger
}

func NewUserService(repo repository.User, logger *slog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) Create(passport serializers.Passport) (uuid.UUID, error) {
	u.logger.Debug("Creating user", "passport", passport)

	if err := passport.Validate(); err != nil {
		u.logger.Debug("Passport validation failed", "error", err)
		return uuid.Nil, err
	}

	u.logger.Debug("Passport validated successfully")

	passportNumber, err := passport.Number()
	if err != nil {
		u.logger.Debug("Failed to get passport number", "error", err)
		return uuid.Nil, err
	}

	passportSeries, err := passport.Series()
	if err != nil {
		u.logger.Debug("Failed to get passport series", "error", err)
		return uuid.Nil, err
	}

	user := &models.User{
		Id:             uuid.New(),
		PassportNumber: passportNumber,
		PassportSeries: passportSeries,
	}

	userId, err := u.repo.Create(user)
	if err != nil {
		u.logger.Error("Failed to create user in repository", "error", err)
		return uuid.Nil, err
	}

	u.logger.Info("User created in repository successfully", "userId", userId)

	return userId, nil
}

func (u *UserService) Get(filters serializers.GetUsersRequest) (*serializers.GetUsersResponse, error) {
	u.logger.Debug("Getting users with filters", "filters", filters)

	users, err := u.repo.Get(filters)
	if err != nil {
		u.logger.Error("Error getting users from repository", "error", err)
		return nil, fmt.Errorf("error getting users: %v", err)
	}

	u.logger.Debug("Users retrieved from repository successfully")

	var responseUsers []serializers.User
	for _, user := range *users {
		responseUsers = append(responseUsers, serializers.User{
			Id:             user.Id,
			PassportNumber: user.PassportNumber,
			PassportSeries: user.PassportSeries,
			Name:           user.Name,
			Surname:        user.Surname,
			Patronymic:     user.Patronymic,
			Address:        user.Address,
		})
	}

	response := &serializers.GetUsersResponse{
		Info: responseUsers,
	}

	u.logger.Info("GetUsersResponse formed successfully", "response", response)

	return response, nil
}

func (u *UserService) Update(userId uuid.UUID, updateInfo serializers.UpdateUserRequest) error {
	u.logger.Debug("Updating user", "userId", userId, "updateInfo", updateInfo)

	updatedUser := &models.User{}
	if updateInfo.PassportNumber != nil {
		updatedUser.PassportNumber = string(*updateInfo.PassportNumber)
	}

	if updateInfo.PassportSeries != nil {
		updatedUser.PassportSeries = string(*updateInfo.PassportSeries)
	}

	if updateInfo.Name != nil {
		updatedUser.Name = *updateInfo.Name
	}

	if updateInfo.Surname != nil {
		updatedUser.Surname = *updateInfo.Surname
	}

	if updateInfo.Patronymic != nil {
		updatedUser.Patronymic = *updateInfo.Patronymic
	}

	if updateInfo.Address != nil {
		updatedUser.Address = *updateInfo.Address
	}

	if err := u.repo.Update(userId, updatedUser); err != nil {
		u.logger.Error("Failed to update user in repository", "error", err)
		return fmt.Errorf("error updating user: %v", err)
	}

	u.logger.Info("User updated in repository successfully", "userId", userId)

	return nil
}

func (u *UserService) Delete(ctx context.Context, userId uuid.UUID) error {
	u.logger.Debug("Deleting user", "userId", userId)

	if err := u.repo.Delete(ctx, userId); err != nil {
		u.logger.Error("Failed to delete user from repository", "error", err)
		return fmt.Errorf("error deleting user: %v", err)
	}

	u.logger.Info("User deleted from repository successfully", "userId", userId)

	return nil
}
