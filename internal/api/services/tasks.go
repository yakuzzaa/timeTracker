package services

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/yakuzzaa/timeTracker/internal/api/repository"
	"github.com/yakuzzaa/timeTracker/internal/api/serializers"
	"github.com/yakuzzaa/timeTracker/internal/models"
)

type TaskService struct {
	repo   repository.Task
	logger *slog.Logger
}

func NewTaskService(repo repository.Task, logger *slog.Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

func (t *TaskService) Create(userId uuid.UUID) (uuid.UUID, error) {
	t.logger.Debug("Creating task for user", "userId", userId)

	task := &models.Task{
		Id:     uuid.New(),
		UserId: userId,
	}

	taskId, err := t.repo.Create(task)
	if err != nil {
		t.logger.Error("Failed to create task in repository", "error", err)
		return uuid.Nil, fmt.Errorf("error creating task: %v", err)
	}

	t.logger.Info("Task created in repository successfully", "taskId", taskId)

	return taskId, nil
}

func (t *TaskService) Update(userId uuid.UUID, taskId uuid.UUID) error {
	t.logger.Debug("Updating task", "userId", userId, "taskId", taskId)

	if err := t.repo.Update(userId, taskId); err != nil {
		t.logger.Error("Failed to update task in repository", "error", err)
		return fmt.Errorf("error updating task: %v", err)
	}

	t.logger.Info("Task updated in repository successfully", "taskId", taskId)

	return nil
}

func (t *TaskService) Get(userId uuid.UUID) (*serializers.GetTaskResponse, error) {
	t.logger.Debug("Getting tasks for user", "userId", userId)

	tasks, err := t.repo.Get(userId)
	if err != nil {
		t.logger.Error("Error getting tasks from repository", "error", err)
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}

	t.logger.Debug("Tasks retrieved from repository successfully")

	var responseTasks []serializers.Task
	for _, task := range *tasks {
		responseTasks = append(responseTasks, serializers.Task{
			Id:        task.Id,
			UserID:    task.UserId,
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
			Total:     t.formatTotalTime(task.Total),
		})
	}

	response := &serializers.GetTaskResponse{
		Info: responseTasks,
	}

	t.logger.Info("GetTaskResponse formed successfully", "response", response)

	return response, nil
}

func (t *TaskService) formatTotalTime(total string) string {
	parts := strings.Split(total, ":")
	hours := parts[0]
	minutes := parts[1]

	return hours + ":" + minutes
}
