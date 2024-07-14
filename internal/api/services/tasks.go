package services

import (
	"fmt"
	"strings"
	"timeTracker/internal/api/repository"
	"timeTracker/internal/api/serializers"
	"timeTracker/internal/storage/models"

	"github.com/google/uuid"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(userId uuid.UUID) (uuid.UUID, error) {
	task := &models.Task{
		Id:     uuid.New(),
		UserId: userId,
	}

	return t.repo.Create(task)
}

func (t *TaskService) Update(userId uuid.UUID, taskId uuid.UUID) error {
	return t.repo.Update(userId, taskId)
}

func (t *TaskService) Get(userId uuid.UUID) (*serializers.GetTaskResponse, error) {
	tasks, err := t.repo.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}

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

	return response, nil
}

func (t *TaskService) formatTotalTime(total string) string {
	parts := strings.Split(total, ":")
	hours := parts[0]
	minutes := parts[1]

	return hours + ":" + minutes
}
