package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// CommonFileGenerator handles generation of common project files
type CommonFileGenerator struct {
	config *Config
	logger Logger
}

// NewCommonFileGenerator creates a new common file generator
func NewCommonFileGenerator(config *Config, logger Logger) *CommonFileGenerator {
	return &CommonFileGenerator{
		config: config,
		logger: logger,
	}
}

// GenerateGoMod generates go.mod file
func (cfg *CommonFileGenerator) GenerateGoMod(projectPath string) error {
	content := fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.9
)`, cfg.config.ProjectName)

	filePath := filepath.Join(projectPath, "go.mod")
	return os.WriteFile(filePath, []byte(content), 0644)
}

// GenerateReadme generates README.md file
func (cfg *CommonFileGenerator) GenerateReadme(projectPath string) error {
	content := fmt.Sprintf(`# %s

A Go application built with %s architecture.

## Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL (optional)

### Installation

1. Clone the repository
2. Install dependencies:
`+"```bash"+`
go mod tidy
`+"```"+`

3. Run the application:
`+"```bash"+`
make run
`+"```"+`

## Architecture

This project follows the %s architecture pattern.

## API Endpoints

- `+"`GET /health`"+` - Health check
- `+"`GET /api/v1/users`"+` - Get all users
- `+"`POST /api/v1/users`"+` - Create a new user

## Development

### Running Tests
`+"```bash"+`
make test
`+"```"+`

### Building
`+"```bash"+`
make build
`+"```"+`

### Docker
`+"```bash"+`
make docker-build
make docker-run
`+"```"+`

## License

This project is licensed under the %s License.`,
		cfg.config.ProjectName,
		cfg.config.Architecture,
		cfg.config.Architecture,
		cfg.config.License)

	filePath := filepath.Join(projectPath, "README.md")
	return os.WriteFile(filePath, []byte(content), 0644)
}

// GenerateGitignore generates .gitignore file
func (cfg *CommonFileGenerator) GenerateGitignore(projectPath string) error {
	content := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool
*.out
*.cover

# Dependency directories
vendor/

# Go workspace file
go.work
go.work.sum

# Build artifacts
/bin/
/dist/
/build/

# Environment variables
.env
.env.local
.env.*.local

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS generated files
.DS_Store
Thumbs.db

# Logs
*.log
logs/

# Database
*.db
*.sqlite
*.sqlite3

# Temporary files
tmp/
temp/`

	filePath := filepath.Join(projectPath, ".gitignore")
	return os.WriteFile(filePath, []byte(content), 0644)
}
