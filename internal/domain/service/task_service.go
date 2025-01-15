package service

import (
	"to-do-list/internal/domain/models"
	"to-do-list/internal/domain/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.repo.Create(task)
}
