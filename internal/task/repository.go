package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(task *Task) error
	Update(task *Task) error
	FindByID(id uint) (*Task, error)
	Delete(task *Task) error
	FindAllByUserID(userID uint) ([]Task, error)
}

type repository struct {
	db *gorm.DB
}

// Create implements Repository.
func (r *repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

// Delete implements Repository.
func (r *repository) Delete(task *Task) error {
	return r.db.Delete(task).Error
}

// FindAllByUserID implements Repository.
func (r *repository) FindAllByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	if err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// FindByID implements Repository.
func (r *repository) FindByID(id uint) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// Update implements Repository.
func (r *repository) Update(task *Task) error {
	return r.db.Save(task).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}