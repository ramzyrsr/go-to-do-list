package repository

import (
	"database/sql"
	"to-do-list/internal/domain/models"
)

type TaskRepository interface {
	Create(task *models.Task) (*models.Task, error)
	// GetAllTask(task *models.Task) (*models.Task, error)
	// Update(id int) (string, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) (*models.Task, error) {
	query := "INSERT INTO tasks(user_id, data) VALUES($1, $2)"
	err := r.db.QueryRow(query, task.UserId, task.Data).Err()
	if err != nil {
		return nil, err
	}

	return task, nil
}
