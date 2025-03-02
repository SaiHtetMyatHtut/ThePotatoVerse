package handler

import (
	"net/http"
	"time"

	"github.com/ThePotatoVerse/internal/app/graphql"
	"github.com/ThePotatoVerse/internal/app/middleware"
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
	authService := service.NewAuthService(log, userRepo, "your-jwt-secret", 24*time.Hour)

	// Create middleware
	authMiddleware := middleware.AuthMiddleware(authService, log)

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		authHandler := NewAuthHandler(authService, log)
		authHandler.RegisterRoutes(api)

		// User routes (protected by auth)
		userHandler := NewUserHandler(log, userService)
		users := api.Group("/users")
		users.Use(authMiddleware)
		{
			users.GET("", userHandler.List)
			users.GET("/:id", userHandler.Get)
			users.POST("", userHandler.Create)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	// GraphQL endpoint
	graphqlServer, err := graphql.NewServer(log, authService)
	if err != nil {
		log.Fatal("Failed to create GraphQL server", "error", err)
	}
	router.POST("/graphql", gin.WrapH(graphqlServer.Handler()))

	return router
}

// loggerMiddleware creates a middleware for logging requests
func loggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		log.Info("Request",
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
		)
	}
}
