package serializers

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	UserID    uuid.UUID
	StartTime time.Time
	EndTime   time.Time
	Total     string
}

type CreateTaskResponse struct {
	Id     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type UpdateTaskRequest struct {
	Id uuid.UUID `json:"task_id"`
}
type UpdateTaskResponse struct {
	Status string
}

type GetTaskResponse struct {
	Info []Task `json:"info"`
}
