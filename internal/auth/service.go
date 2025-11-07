package auth

import (
	"errors"
	"rest-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Claims struct for JWT payload
type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

type Service interface {
	Register(username, email, password string) (*UserResponse, error)
	Login(email, password string) (string, *UserResponse, error)
	GenerateToken(userID uint) (string, error)
}

type service struct {
	repo Repository
	cfg  *config.Config
}

// GenerateToken implements Service.
func (s *service) GenerateToken(userID uint) (string, error) {
	// Parse duration from config
	duration, err := time.ParseDuration(s.cfg.JWTExpires)
	if err != nil {
		duration = 168 * time.Hour // Default 7 days
	}

	// Create claims with user ID and standard claims
	claims := Claims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with signing method HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Sign token with secret key and return token string
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

// Login implements Service.
func (s *service) Login(email string, password string) (string, *UserResponse, error) {
	if email == "" || password == "" {
		return "", nil, errors.New("email dan password tidak boleh kosong")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("email atau password salah")
		}
		return "", nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	// Generate token
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	userResponse := &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return token, userResponse, nil
}

// Register implements Service.
func (s *service) Register(username string, email string, password string) (*UserResponse, error) {
	// Validasi input
	if username == "" || email == "" || password == "" {
		return nil, errors.New("username, email, dan password tidak boleh kosong")
	}

	// Check if user already exists
	existingUser, err := s.repo.FindEmailOrUsername(email, username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email atau username sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Register(user); err != nil {
		return nil, err
	}

	userResponse := &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}

func NewService(repo Repository, cfg *config.Config) Service {
	return &service{repo: repo, cfg: cfg}
}

func (s *service) GetTokenExpiration() time.Duration {
	duration, err := time.ParseDuration(s.cfg.JWTExpires)
	if err != nil {
		return 7 * 24 * time.Hour // default 7 days
	}
	return duration
}