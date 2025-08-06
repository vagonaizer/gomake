package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gomake/pkg/logger"
)

// Config holds the configuration for project generation
type Config struct {
	ProjectName  string
	Architecture string
	TargetDir    string
	WithDocker   bool
	WithMakefile bool
	WithGit      bool
	License      string
	AutoYes      bool
}

// Generator handles project generation
type Generator struct {
	config  *Config
	logger  *logger.Logger
	arch    Architecture
	fileGen *FileGenerator
}

// Architecture interface defines methods for different architectures
type Architecture interface {
	GetName() string
	GetStructure() *ProjectStructure
	GenerateFiles(projectPath string, config *Config) error
}

// New creates a new generator instance
func New(config *Config, logger *logger.Logger) (*Generator, error) {
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}

	// Create architecture instance
	arch, err := createArchitecture(config.Architecture)
	if err != nil {
		return nil, fmt.Errorf("failed to create architecture: %w", err)
	}

	// Create file generator
	fileGen, err := NewFileGenerator(config, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create file generator: %w", err)
	}

	return &Generator{
		config:  config,
		logger:  logger,
		arch:    arch,
		fileGen: fileGen,
	}, nil
}

// Generate creates the project structure
func (g *Generator) Generate() error {
	projectPath := filepath.Join(g.config.TargetDir, g.config.ProjectName)

	g.logger.Info("Creating project directory", "path", projectPath)

	// Create project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Generate directory structure
	if err := g.generateStructure(projectPath); err != nil {
		return fmt.Errorf("failed to generate structure: %w", err)
	}

	// Generate architecture-specific files
	if err := g.arch.GenerateFiles(projectPath, g.config); err != nil {
		return fmt.Errorf("failed to generate architecture files: %w", err)
	}

	// Generate common files using FileGenerator
	if err := g.fileGen.GenerateCommonFiles(projectPath); err != nil {
		return fmt.Errorf("failed to generate common files: %w", err)
	}

	// Generate optional files using FileGenerator
	if err := g.fileGen.GenerateOptionalFiles(projectPath); err != nil {
		return fmt.Errorf("failed to generate optional files: %w", err)
	}

	g.logger.Success("Project generated successfully", "path", projectPath)
	return nil
}

func (g *Generator) generateStructure(projectPath string) error {
	structure := g.arch.GetStructure()

	for _, dir := range structure.Directories {
		dirPath := filepath.Join(projectPath, dir)
		g.logger.Debug("Creating directory", "path", dirPath)

		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func createArchitecture(archType string) (Architecture, error) {
	switch archType {
	case "hexagonal":
		return NewHexagonalArchitecture(), nil
	case "clean":
		return NewCleanArchitecture(), nil
	case "mvc":
		return NewMVCArchitecture(), nil
	case "basic":
		return NewBasicArchitecture(), nil
	default:
		return nil, fmt.Errorf("unsupported architecture: %s", archType)
	}
}
