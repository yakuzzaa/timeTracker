package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/yakuzzaa/timeTracker/internal/api/serializers"
	models2 "github.com/yakuzzaa/timeTracker/internal/models"
	"gorm.io/gorm"
)

type User interface {
	Create(user *models2.User) (uuid.UUID, error)
	Get(filters serializers.GetUsersRequest) (*[]models2.User, error)
	Update(userId uuid.UUID, user *models2.User) error
	Delete(ctx context.Context, userId uuid.UUID) error
}

type Task interface {
	Create(task *models2.Task) (uuid.UUID, error)
	Update(userId uuid.UUID, taskId uuid.UUID) error
	Get(userId uuid.UUID) (*[]models2.Task, error)
}

type Repository struct {
	User
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Task: NewTaskRepository(db),
	}
}
