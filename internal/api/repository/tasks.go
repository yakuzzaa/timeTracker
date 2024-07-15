package repository

import (
	"fmt"

	"github.com/yakuzzaa/timeTracker/internal/storage/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(task *models.Task) (uuid.UUID, error) {
	var taskId uuid.UUID
	query := `
		INSERT INTO tasks (id, user_id)
		VALUES (?, ?)
		RETURNING id
	`

	row := t.db.Raw(query, task.Id, task.UserId).Row()

	if err := row.Scan(&taskId); err != nil {
		return uuid.Nil, err
	}

	return taskId, nil
}

func (t *TaskRepository) Update(userId uuid.UUID, taskId uuid.UUID) error {
	if err := t.isUserTaskExist(userId, taskId); err != nil {
		return err
	}

	updateQuery := `UPDATE tasks SET end_time = NOW() WHERE id = ?`
	if err := t.db.Exec(updateQuery, taskId).Error; err != nil {
		return fmt.Errorf("error updating task: %v", err)
	}

	return nil

}

func (t *TaskRepository) Get(userId uuid.UUID) (*[]models.Task, error) {
	var tasks []models.Task

	query := `SELECT * FROM tasks WHERE user_id = ? ORDER BY total DESC`
	if err := t.db.Raw(query, userId).Scan(&tasks).Error; err != nil {
		return nil, fmt.Errorf("error fetching tasks for user %v: %v", userId, err)
	}

	return &tasks, nil
}

func (t *TaskRepository) isUserTaskExist(userId, taskId uuid.UUID) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ? AND user_id = ?)`
	if err := t.db.Raw(checkQuery, taskId, userId).Scan(&exists).Error; err != nil {
		return fmt.Errorf("error checking task existence for user: %v", err)
	}

	if !exists {
		return fmt.Errorf("task with ID %v not found for user with ID %v", taskId, userId)
	}
	return nil
}
