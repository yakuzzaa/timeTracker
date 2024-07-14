package services

import (
	"context"
	"timeTracker/internal/api/repository"
	"timeTracker/internal/api/serializers"
	"timeTracker/internal/storage/models"

	"github.com/google/uuid"
)

type User interface {
	Create(passportNumber serializers.Passport) (uuid.UUID, error)
	Get(filters serializers.GetUsersRequest) (*serializers.GetUsersResponse, error)
	Update(userId uuid.UUID, updateInfo serializers.UpdateUserRequest) error
	Delete(ctx context.Context, userId uuid.UUID) error
}

type Task interface {
	Create(userId uuid.UUID) (models.Task, error)
	Update(userId uuid.UUID, taskId uuid.UUID) (models.Task, error)
	Get(userId uuid.UUID) (models.Task, error)
}

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		//Task: NewTaskService(repos.Task),
	}
}
