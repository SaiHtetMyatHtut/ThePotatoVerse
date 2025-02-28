package handler

import (
	"net/http"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/service"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	log         logger.Logger
	userService service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(log logger.Logger, userService service.UserService) *UserHandler {
	return &UserHandler{
		log:         log,
		userService: userService,
	}
}

// List returns a list of users
func (h *UserHandler) List(c *gin.Context) {
	h.log.Info("Handling list users request")

	users, err := h.userService.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Create creates a new user
func (h *UserHandler) Create(c *gin.Context) {
	h.log.Info("Handling create user request")

	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	createdUser, err := h.userService.Create(c.Request.Context(), user)
	if err != nil {
		if err == service.ErrInvalidInput {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// Get returns a user by ID
func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	h.log.Info("Handling get user request", "id", id)

	user, err := h.userService.Get(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update updates a user
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	h.log.Info("Handling update user request", "id", id)

	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}

	err := h.userService.Update(c.Request.Context(), user)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case service.ErrInvalidInput:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		}
		return
	}

	c.Status(http.StatusOK)
}

// Delete deletes a user
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	h.log.Info("Handling delete user request", "id", id)

	err := h.userService.Delete(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
