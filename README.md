# ThePotatoVerse

A Go application with a clean architecture design.

## Project Structure

```
.
├── api/                  # API documentation
├── cmd/                  # Application entry points
│   └── app/              # Main application
├── config/               # Configuration files
├── internal/             # Private application code
│   ├── app/              # Application core
│   │   ├── handler/      # HTTP handlers
│   │   ├── model/        # Domain models
│   │   ├── repository/   # Data access interfaces and implementations
│   │   └── service/      # Business logic
│   └── pkg/              # Private packages
│       ├── config/       # Configuration loading
│       ├── middleware/   # HTTP middleware
│       └── validator/    # Input validation
├── pkg/                  # Public packages
│   ├── database/         # Database utilities
│   └── logger/           # Logging utilities
├── scripts/              # Scripts and tools
│   └── migrations/       # Database migrations
└── test/                 # Test utilities and mocks
```

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Docker and Docker Compose (for local development)
- Make

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/ThePotatoVerse.git
cd ThePotatoVerse
```

2. Install dependencies:

```bash
make deps
```

3. Start the development environment:

```bash
docker-compose up -d
```

4. Run database migrations:

```bash
make migrate-up
```

5. Run the application:

```bash
make run
```

## Development

### Available Commands

Run `make help` to see all available commands:

```
make build          - Build the application
make run            - Run the application
make test           - Run tests
make test-coverage  - Run tests with coverage
make lint           - Run linter
make clean          - Clean build artifacts
make deps           - Install dependencies
make migrate-up     - Run database migrations up
make migrate-down   - Run database migrations down
make migrate-create - Create a new migration file
make docs           - Generate API documentation
make dev            - Run the application in development mode
```

## API Documentation

API documentation is available at `/swagger/index.html` when the application is running.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 