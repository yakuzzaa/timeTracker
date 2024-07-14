package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	StartTime time.Time
	EndTime   time.Time
	Total     string
}
