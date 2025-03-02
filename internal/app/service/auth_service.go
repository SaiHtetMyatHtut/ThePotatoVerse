package service

import (
	"context"
	"errors"
	"time"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/repository"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Auth-related errors
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidToken       = errors.New("invalid token")
)

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Register(ctx context.Context, input model.RegisterUserInput) (model.AuthResponse, error)
	Login(ctx context.Context, credentials model.UserCredentials) (model.AuthResponse, error)
	VerifyToken(ctx context.Context, token string) (*JWTClaims, error)
}

// authService implements AuthService
type authService struct {
	log       logger.Logger
	userRepo  repository.UserRepository
	jwtSecret string
	jwtExpiry time.Duration
}

// NewAuthService creates a new authentication service
func NewAuthService(
	log logger.Logger,
	userRepo repository.UserRepository,
	jwtSecret string,
	jwtExpiry time.Duration,
) AuthService {
	return &authService{
		log:       log,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
	}
}

// Register creates a new user account
func (s *authService) Register(ctx context.Context, input model.RegisterUserInput) (model.AuthResponse, error) {
	s.log.Info("Registering new user", "email", input.Email)

	// Check if email already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return model.AuthResponse{}, err
	}
	if exists {
		return model.AuthResponse{}, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.AuthResponse{}, err
	}

	// Create user
	now := time.Now()
	user := model.User{
		Name:          input.Name,
		Email:         input.Email,
		PasswordHash:  string(hashedPassword),
		Role:          "user", // Default role
		EmailVerified: false,  // Requires verification
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return model.AuthResponse{}, err
	}

	// Generate token
	token, err := s.generateToken(createdUser)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		User:  createdUser,
		Token: token,
	}, nil
}

// Login authenticates a user and returns a token
func (s *authService) Login(ctx context.Context, credentials model.UserCredentials) (model.AuthResponse, error) {
	s.log.Info("User login attempt", "email", credentials.Email)

	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return model.AuthResponse{}, ErrInvalidCredentials
		}
		return model.AuthResponse{}, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
	if err != nil {
		return model.AuthResponse{}, ErrInvalidCredentials
	}

	// Update last login time
	user.LastLoginAt = time.Now()
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		s.log.Error("Failed to update last login time", "error", err)
		// Continue anyway, this is not critical
	}

	// Generate token
	token, err := s.generateToken(user)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// VerifyToken validates a JWT token and returns the claims
func (s *authService) VerifyToken(ctx context.Context, tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// generateToken creates a new JWT token for a user
func (s *authService) generateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(s.jwtExpiry)

	claims := &JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
