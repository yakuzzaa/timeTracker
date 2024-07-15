package repository

import (
	"context"

	"github.com/yakuzzaa/timeTracker/internal/api/serializers"
	"github.com/yakuzzaa/timeTracker/internal/storage/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	Create(user *models.User) (uuid.UUID, error)
	Get(filters serializers.GetUsersRequest) (*[]models.User, error)
	Update(userId uuid.UUID, user *models.User) error
	Delete(ctx context.Context, userId uuid.UUID) error
}

type Task interface {
	Create(task *models.Task) (uuid.UUID, error)
	Update(userId uuid.UUID, taskId uuid.UUID) error
	Get(userId uuid.UUID) (*[]models.Task, error)
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
