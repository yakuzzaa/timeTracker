package services

import (
	"context"
	"fmt"
	"timeTracker/internal/api/repository"
	"timeTracker/internal/api/serializers"
	"timeTracker/internal/storage/models"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(passport serializers.Passport) (uuid.UUID, error) {
	passportNumber, err := passport.Number()
	if err != nil {
		return uuid.Nil, err
	}
	passportSeries, err := passport.Series()
	if err != nil {
		return uuid.Nil, err
	}
	user := &models.User{
		Id:             uuid.New(),
		PassportNumber: passportNumber,
		PassportSeries: passportSeries,
	}

	return u.repo.Create(user)
}

func (u *UserService) Get(filters serializers.GetUsersRequest) (*serializers.GetUsersResponse, error) {
	users, err := u.repo.Get(filters)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}

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

	return response, nil
}

func (u *UserService) Update(userId uuid.UUID, updateInfo serializers.UpdateUserRequest) error {
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
	return u.repo.Update(userId, updatedUser)
}

func (u *UserService) Delete(ctx context.Context, userId uuid.UUID) error {
	return u.repo.Delete(ctx, userId)
}
