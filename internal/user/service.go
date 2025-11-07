package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	GetUserByID(id uint) (*Response, error)
	UpdateUser(currentUserID, targetUserID uint, req *UpdateRequest) (*Response, error)
	GetProfile(userID uint) (*Response, error)
}

type service struct {
	repo Repository
}

// GetProfile implements Service.
func (s *service) GetProfile(userID uint) (*Response, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get profile")
	}

	response := &Response{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

// GetUserByID implements Service.
func (s *service) GetUserByID(id uint) (*Response, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	response := &Response{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

// UpdateUser implements Service.
func (s *service) UpdateUser(currentUserID, targetUserID uint, req *UpdateRequest) (*Response, error) {
	// Check authorization
	if currentUserID != targetUserID {
		return nil, errors.New("unauthorized to update this user")
	}

	user, err := s.repo.FindByID(targetUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	// Update email if provided
	if req.Email != nil && *req.Email != user.Email {
		exists, err := s.repo.ExistsByEmail(*req.Email)
		if err != nil {
			return nil, errors.New("failed to check email availability")
		}
		if exists {
			return nil, errors.New("email already in use")
		}
		user.Email = *req.Email
	}

	// Update username if provided
	if req.Username != nil && *req.Username != user.Username {
		exists, err := s.repo.ExistsByUsername(*req.Username)
		if err != nil {
			return nil, errors.New("failed to check username availability")
		}
		if exists {
			return nil, errors.New("username already in use")
		}
		user.Username = *req.Username
	}

	// Update password if provided
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), 10)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	}

	// Save changes
	if err := s.repo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	response := &Response{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}