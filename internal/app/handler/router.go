package handler

import (
	"net/http"
	"time"

	"github.com/ThePotatoVerse/internal/app/repository/memory"
	"github.com/ThePotatoVerse/internal/app/service"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a new router
func NewRouter(log logger.Logger) http.Handler {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	router := gin.New()

	// Add middleware
	router.Use(
		gin.Recovery(),
		loggerMiddleware(log),
	)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Initialize repositories
	userRepo := memory.NewUserRepository()

	// Initialize services
	userService := service.NewUserService(log, userRepo)

	// API routes
	api := router.Group("/api/v1")
	{
		// User routes
		userHandler := NewUserHandler(log, userService)
		users := api.Group("/users")
		{
			users.GET("", userHandler.List)
			users.POST("", userHandler.Create)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	return router
}

// loggerMiddleware creates a gin middleware for logging requests
func loggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		log.Info("Request",
			"status", status,
			"method", method,
			"path", path,
			"ip", ip,
			"latency", latency,
		)
	}
}
