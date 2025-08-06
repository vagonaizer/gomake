package generator

import (
	"fmt"
)

// FileGenerator handles generation of common project files
type FileGenerator struct {
	config      *Config
	logger      Logger
	commonGen   *CommonFileGenerator
	makefileGen *MakefileGenerator
	dockerGen   *DockerGenerator
	licenseGen  *LicenseGenerator
	gitGen      *GitGenerator
}

// Logger interface for file generator
type Logger interface {
	Info(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Success(msg string, args ...interface{})
}

// NewFileGenerator creates a new file generator
func NewFileGenerator(config *Config, logger Logger) (*FileGenerator, error) {
	return &FileGenerator{
		config:      config,
		logger:      logger,
		commonGen:   NewCommonFileGenerator(config, logger),
		makefileGen: NewMakefileGenerator(config, logger),
		dockerGen:   NewDockerGenerator(config, logger),
		licenseGen:  NewLicenseGenerator(config, logger),
		gitGen:      NewGitGenerator(config, logger),
	}, nil
}

// GenerateCommonFiles generates all common project files
func (fg *FileGenerator) GenerateCommonFiles(projectPath string) error {
	fg.logger.Info("Generating common files")

	if err := fg.commonGen.GenerateGoMod(projectPath); err != nil {
		return fmt.Errorf("failed to generate go.mod: %w", err)
	}

	if err := fg.commonGen.GenerateReadme(projectPath); err != nil {
		return fmt.Errorf("failed to generate README.md: %w", err)
	}

	if err := fg.commonGen.GenerateGitignore(projectPath); err != nil {
		return fmt.Errorf("failed to generate .gitignore: %w", err)
	}

	return nil
}

// GenerateOptionalFiles generates optional files based on configuration
func (fg *FileGenerator) GenerateOptionalFiles(projectPath string) error {
	fg.logger.Info("Generating optional files")

	// Always generate Makefile for better DX
	if err := fg.makefileGen.Generate(projectPath); err != nil {
		return fmt.Errorf("failed to generate Makefile: %w", err)
	}

	// Generate Docker files if requested
	if fg.config.WithDocker {
		if err := fg.dockerGen.Generate(projectPath); err != nil {
			return fmt.Errorf("failed to generate Docker files: %w", err)
		}
	}

	// Generate license file if specified
	if fg.config.License != "None" && fg.config.License != "" {
		if err := fg.licenseGen.Generate(projectPath); err != nil {
			return fmt.Errorf("failed to generate license: %w", err)
		}
	}

	// Initialize git repository if requested
	if fg.config.WithGit {
		if err := fg.gitGen.Initialize(projectPath); err != nil {
			return fmt.Errorf("failed to initialize git repository: %w", err)
		}
	}

	return nil
}
