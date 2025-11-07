package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*User, error)
	FindEmailOrUsername(email, username string) (*User, error)
	Register(user *User) error
	FindByID(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

// FindByEmail implements Repository.
func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID implements Repository.
func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindEmailOrUsername implements Repository.
func (r *repository) FindEmailOrUsername(email string, username string) (*User, error) {
	var user User
	err := r.db.Where("email = ? OR username = ?", email, username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Register implements Repository.
func (r *repository) Register(user *User) error {
	return r.db.Create(user).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}