package handler

import (
	"net/http"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/service"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authService service.AuthService
	log         logger.Logger
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService service.AuthService, log logger.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         log,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var input model.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Error("Invalid registration input", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := h.authService.Register(c.Request.Context(), input)
	if err != nil {
		h.log.Error("Registration failed", "error", err, "email", input.Email)

		if err == service.ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials model.UserCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		h.log.Error("Invalid login input", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := h.authService.Login(c.Request.Context(), credentials)
	if err != nil {
		h.log.Error("Login failed", "error", err, "email", credentials.Email)

		if err == service.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCurrentUser returns the current authenticated user
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// In a real implementation, you would call a user service to get the user by ID
	// For now, we'll just return a placeholder with the information from the context
	user := model.User{
		ID:    userID.(string),
		Email: c.GetString("email"),
		Role:  c.GetString("role"),
	}

	c.JSON(http.StatusOK, user)
}

// RegisterRoutes registers the auth routes
func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.GET("/me", h.GetCurrentUser) // This route should be protected by auth middleware
	}
}
