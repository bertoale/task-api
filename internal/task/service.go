package task

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateTask(userID uint, title, description string) (*Response, error)
	GetTasksByUserID(userID uint) ([]Response, error)
	GetTaskByID(userID, id uint) (*Response, error)
	UpdateTask(userID, taskID uint, req *UpdateRequest) (*Response, error)
	DeleteTask(userID, taskID uint) error
}

type service struct {
	repo Repository
}

// CreateTask implements Service.
func (s *service) CreateTask(userID uint, title, description string) (*Response, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	task := &Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		IsCompleted: false,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, errors.New("failed to create task")
	}

	response := &Response{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}

	return response, nil
}

// DeleteTask implements Service.
func (s *service) DeleteTask(userID, taskID uint) error {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return errors.New("failed to retrieve task")
	}

	if task.UserID != userID {
		return errors.New("unauthorized to delete this task")
	}

	if err := s.repo.Delete(task); err != nil {
		return errors.New("failed to delete task")
	}

	return nil
}

// GetTaskByID implements Service.
func (s *service) GetTaskByID(userID, id uint) (*Response, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, errors.New("failed to retrieve task")
	}

	if task.UserID != userID {
		return nil, errors.New("unauthorized to access this task")
	}

	response := &Response{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}

	return response, nil
}

// GetTasksByUserID implements Service.
func (s *service) GetTasksByUserID(userID uint) ([]Response, error) {
	tasks, err := s.repo.FindAllByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve tasks")
	}

	if len(tasks) == 0 {
		return []Response{}, nil
	}

	responses := make([]Response, len(tasks))
	for i, task := range tasks {
		responses[i] = Response{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			IsCompleted: task.IsCompleted,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			UserID:      task.UserID,
		}
	}

	return responses, nil
}

// UpdateTask implements Service.
func (s *service) UpdateTask(userID, taskID uint, req *UpdateRequest) (*Response, error) {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, errors.New("failed to retrieve task")
	}

	if task.UserID != userID {
		return nil, errors.New("unauthorized to update this task")
	}

	// Update fields that were provided
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.IsCompleted != nil {
		task.IsCompleted = *req.IsCompleted
	}

	if err := s.repo.Update(task); err != nil {
		return nil, errors.New("failed to update task")
	}

	response := &Response{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}

	return response, nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}