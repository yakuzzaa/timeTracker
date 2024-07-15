package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/yakuzzaa/timeTracker/internal/api/serializers"
	"github.com/yakuzzaa/timeTracker/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.User) (uuid.UUID, error) {
	var userId uuid.UUID

	query := `
		INSERT INTO users (id, passport_number, passport_series)
		VALUES (?, ?, ?)
		RETURNING id
	`

	row := u.db.Raw(query, user.Id, user.PassportNumber, user.PassportSeries).Row()

	if err := row.Scan(&userId); err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (u *UserRepository) Get(filters serializers.GetUsersRequest) (*[]models.User, error) {
	builder := NewUserQueryBuilder()

	builder, err := builder.WithPassportNumber(filters.PassportNumber)
	if err != nil {
		return nil, err
	}
	builder, err = builder.WithPassportSeries(filters.PassportSeries)
	if err != nil {
		return nil, err
	}
	builder, err = builder.
		WithID(filters.Id).
		WithName(filters.Name).
		WithSurname(filters.Surname).
		WithPatronymic(filters.Patronymic).
		WithAddress(filters.Address).
		WithPagination(filters.Page, filters.PageSize)
	if err != nil {
		return nil, err
	}
	sql, args := builder.Build()

	var users []models.User
	if err := u.db.Raw(sql, args...).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepository) Update(userId uuid.UUID, user *models.User) error {
	if err := u.isUserExists(userId); err != nil {
		return err
	}

	query := `
		UPDATE users
		SET passport_number = COALESCE(NULLIF(?, ''), passport_number),
		    passport_series = COALESCE(NULLIF(?, ''), passport_series),
		    surname = COALESCE(NULLIF(?, ''), surname),
		    name = COALESCE(NULLIF(?, ''), name),
		    patronymic = COALESCE(NULLIF(?, ''), patronymic),
		    address = COALESCE(NULLIF(?, ''), address)
		WHERE id = ?
	`

	if err := u.db.Exec(query, user.PassportNumber, user.PassportSeries, user.Surname, user.Name, user.Patronymic, user.Address, userId).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("DELETE FROM tasks WHERE user_id = ?", userId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Exec("DELETE FROM users WHERE id = ?", userId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) isUserExists(userId uuid.UUID) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`
	if err := u.db.Raw(checkQuery, userId).Scan(&exists).Error; err != nil {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if !exists {
		return fmt.Errorf("user with ID %v not found", userId)
	}
	return nil
}
