package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ThePotatoVerse/internal/pkg/config"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Postgres represents a PostgreSQL database connection
type Postgres struct {
	Pool *pgxpool.Pool
	log  logger.Logger
}

// NewPostgres creates a new PostgreSQL connection
func NewPostgres(ctx context.Context, cfg *config.DBConfig, log logger.Logger) (*Postgres, error) {
	log.Info("Connecting to PostgreSQL database", "host", cfg.Host, "port", cfg.Port, "database", cfg.Name)

	// Create connection string
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	// Create connection pool configuration
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// Set connection pool parameters
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	// Create connection pool
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info("Connected to PostgreSQL database")

	return &Postgres{
		Pool: pool,
		log:  log,
	}, nil
}

// Close closes the database connection
func (p *Postgres) Close() {
	p.log.Info("Closing PostgreSQL connection")
	p.Pool.Close()
}
