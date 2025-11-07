package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) error
	ExistsByEmail(email string) (bool, error)
	ExistsByUsername(username string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// ExistsByEmail implements Repository.
func (r *repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByUsername implements Repository.
func (r *repository) ExistsByUsername(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
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

// FindByUsername implements Repository.
func (r *repository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements Repository.
func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}