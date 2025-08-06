package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// TemplateConfig represents a custom template configuration
type TemplateConfig struct {
	Name         string            `yaml:"name"`
	Description  string            `yaml:"description"`
	Directories  []string          `yaml:"directories"`
	Files        map[string]string `yaml:"files"`
	Dependencies []string          `yaml:"dependencies"`
	Variables    map[string]string `yaml:"variables"`
}

// ConfigFile represents the gomake configuration file
type ConfigFile struct {
	Templates []TemplateConfig `yaml:"templates"`
	Defaults  struct {
		Architecture string `yaml:"architecture"`
		License      string `yaml:"license"`
		WithDocker   bool   `yaml:"with_docker"`
		WithMakefile bool   `yaml:"with_makefile"`
		WithGit      bool   `yaml:"with_git"`
	} `yaml:"defaults"`
}

// LoadConfig loads configuration from .gomake.yml file
func LoadConfig() (*ConfigFile, error) {
	configPaths := []string{
		".gomake.yml",
		".gomake.yaml",
		filepath.Join(os.Getenv("HOME"), ".gomake.yml"),
		filepath.Join(os.Getenv("HOME"), ".config", "gomake", "config.yml"),
	}

	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			return loadConfigFromFile(path)
		}
	}

	// Return default config if no file found
	return getDefaultConfig(), nil
}

func loadConfigFromFile(path string) (*ConfigFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var config ConfigFile
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", path, err)
	}

	return &config, nil
}

func getDefaultConfig() *ConfigFile {
	return &ConfigFile{
		Templates: []TemplateConfig{},
		Defaults: struct {
			Architecture string `yaml:"architecture"`
			License      string `yaml:"license"`
			WithDocker   bool   `yaml:"with_docker"`
			WithMakefile bool   `yaml:"with_makefile"`
			WithGit      bool   `yaml:"with_git"`
		}{
			Architecture: "basic",
			License:      "MIT",
			WithDocker:   false,
			WithMakefile: false,
			WithGit:      false,
		},
	}
}

// SaveDefaultConfig creates a default configuration file
func SaveDefaultConfig(path string) error {
	config := getDefaultConfig()

	// Add example custom template
	config.Templates = []TemplateConfig{
		{
			Name:        "microservice",
			Description: "Microservice with gRPC and HTTP",
			Directories: []string{
				"cmd/server",
				"internal/grpc",
				"internal/http",
				"internal/service",
				"internal/repository",
				"proto",
				"migrations",
			},
			Files: map[string]string{
				"cmd/server/main.go": "// Custom microservice main file\npackage main\n\nfunc main() {\n\t// TODO: Implement\n}",
			},
			Dependencies: []string{
				"google.golang.org/grpc",
				"github.com/grpc-ecosystem/grpc-gateway/v2",
			},
			Variables: map[string]string{
				"service_name": "{{.ProjectName}}",
				"port":         "8080",
				"grpc_port":    "9090",
			},
		},
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
