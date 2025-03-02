package resolver

import (
	"context"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/service"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/graph-gophers/graphql-go"
)

// AuthResolver handles authentication-related GraphQL operations
type AuthResolver struct {
	authService service.AuthService
	log         logger.Logger
}

// NewAuthResolver creates a new auth resolver
func NewAuthResolver(authService service.AuthService, log logger.Logger) *AuthResolver {
	return &AuthResolver{
		authService: authService,
		log:         log,
	}
}

// RegisterInput represents the input for user registration
type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginInput represents the input for user login
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	user  *UserResolver
	token string
}

// User returns the user resolver
func (r *AuthResponse) User() *UserResolver {
	return r.user
}

// Token returns the authentication token
func (r *AuthResponse) Token() string {
	return r.token
}

// Register handles user registration
func (r *AuthResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	registerInput := model.RegisterUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	authResponse, err := r.authService.Register(ctx, registerInput)
	if err != nil {
		r.log.Error("Registration failed", "error", err, "email", input.Email)
		return nil, err
	}

	return &AuthResponse{
		user:  &UserResolver{user: authResponse.User},
		token: authResponse.Token,
	}, nil
}

// Login handles user login
func (r *AuthResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	credentials := model.UserCredentials{
		Email:    input.Email,
		Password: input.Password,
	}

	authResponse, err := r.authService.Login(ctx, credentials)
	if err != nil {
		r.log.Error("Login failed", "error", err, "email", input.Email)
		return nil, err
	}

	return &AuthResponse{
		user:  &UserResolver{user: authResponse.User},
		token: authResponse.Token,
	}, nil
}

// Me returns the currently authenticated user
func (r *AuthResolver) Me(ctx context.Context) (*UserResolver, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, nil // Return null for unauthenticated users
	}

	// Get user from user service (assuming it's available)
	// This would typically be implemented by calling a user service
	// For now, we'll just return a placeholder
	// In a real implementation, you would inject the user service and call it here

	// Example:
	// user, err := r.userService.GetUserByID(ctx, userID)
	// if err != nil {
	//     r.log.Error("Failed to get user", "error", err, "user_id", userID)
	//     return nil, err
	// }
	// return &UserResolver{user: user}, nil

	// Placeholder implementation
	return &UserResolver{user: model.User{
		ID:    userID,
		Email: ctx.Value("email").(string),
		Role:  ctx.Value("role").(string),
	}}, nil
}

// UserResolver resolves User fields
type UserResolver struct {
	user model.User
}

// ID returns the user ID
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

// Name returns the user name
func (r *UserResolver) Name() string {
	return r.user.Name
}

// Email returns the user email
func (r *UserResolver) Email() string {
	return r.user.Email
}

// Role returns the user role
func (r *UserResolver) Role() string {
	return r.user.Role
}

// EmailVerified returns whether the user's email is verified
func (r *UserResolver) EmailVerified() bool {
	return r.user.EmailVerified
}

// LastLoginAt returns the user's last login time
func (r *UserResolver) LastLoginAt() *string {
	if r.user.LastLoginAt.IsZero() {
		return nil
	}
	lastLogin := r.user.LastLoginAt.String()
	return &lastLogin
}

// CreatedAt returns the user creation time
func (r *UserResolver) CreatedAt() string {
	return r.user.CreatedAt.String()
}

// UpdatedAt returns the user update time
func (r *UserResolver) UpdatedAt() string {
	return r.user.UpdatedAt.String()
}
