package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// DockerGenerator handles Docker files generation
type DockerGenerator struct {
	config *Config
	logger Logger
}

// NewDockerGenerator creates a new Docker generator
func NewDockerGenerator(config *Config, logger Logger) *DockerGenerator {
	return &DockerGenerator{
		config: config,
		logger: logger,
	}
}

// Generate creates Docker files
func (dg *DockerGenerator) Generate(projectPath string) error {
	dg.logger.Info("Generating Docker files")

	if err := dg.generateDockerfile(projectPath); err != nil {
		return fmt.Errorf("failed to generate Dockerfile: %w", err)
	}

	if err := dg.generateDockerCompose(projectPath); err != nil {
		return fmt.Errorf("failed to generate docker-compose.yml: %w", err)
	}

	if err := dg.generateDockerignore(projectPath); err != nil {
		return fmt.Errorf("failed to generate .dockerignore: %w", err)
	}

	return nil
}

func (dg *DockerGenerator) generateDockerfile(projectPath string) error {
	content := fmt.Sprintf(`# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/%s/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./main"]
`, dg.config.ProjectName)

	filePath := filepath.Join(projectPath, "Dockerfile")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func (dg *DockerGenerator) generateDockerCompose(projectPath string) error {
	content := fmt.Sprintf(`version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=development
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=password
      - DB_NAME=%s_db
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    depends_on:
      - postgres
      - redis
    restart: unless-stopped

  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=%s_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
`, dg.config.ProjectName, dg.config.ProjectName)

	filePath := filepath.Join(projectPath, "docker-compose.yml")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func (dg *DockerGenerator) generateDockerignore(projectPath string) error {
	content := `# Git
.git
.gitignore

# Documentation
README.md
CHANGELOG.md
LICENSE

# Development files
.env
.env.local
.env.*.local

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS files
.DS_Store
Thumbs.db

# Build artifacts
bin/
dist/
build/

# Test files
coverage.out
coverage.html

# Temporary files
tmp/
temp/

# Node modules (if any)
node_modules/

# Logs
*.log
logs/
`

	filePath := filepath.Join(projectPath, ".dockerignore")
	return os.WriteFile(filePath, []byte(content), 0644)
}
