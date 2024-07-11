package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID `json:"id" db:"id"`
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
	Total     string    `json:"total" db:"total"`
}
