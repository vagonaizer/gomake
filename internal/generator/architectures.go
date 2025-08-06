package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// HexagonalArchitecture implements hexagonal (ports & adapters) architecture
type HexagonalArchitecture struct {
	templateManager *TemplateManager
}

func NewHexagonalArchitecture() *HexagonalArchitecture {
	tm, err := NewTemplateManager()
	if err != nil {
		panic(fmt.Sprintf("Failed to create template manager: %v", err))
	}
	
	return &HexagonalArchitecture{
		templateManager: tm,
	}
}

func (h *HexagonalArchitecture) GetName() string {
	return "hexagonal"
}

func (h *HexagonalArchitecture) GetStructure() *ProjectStructure {
	structure := NewProjectStructure()
	
	// Add directories based on your specification
	dirs := []string{
		"cmd",
		"images",
		"internal/adapters/cache",
		"internal/adapters/handler",
		"internal/adapters/repository",
		"internal/adapters/tests/integration",
		"internal/adapters/tests/unit",
		"internal/config",
		"internal/core/domain",
		"internal/core/ports",
		"internal/core/services",
		"internal/web",
		"pkg/logger",
		"pkg/utils",
		"pkg/database",
		"pkg/initializers",
	}
	
	for _, dir := range dirs {
		structure.AddDirectory(dir)
	}
	
	return structure
}

func (h *HexagonalArchitecture) GenerateFiles(projectPath string, config *Config) error {
	templateData := NewTemplateData(config)
	
	files := map[string]string{
		fmt.Sprintf("cmd/%s/main.go", config.ProjectName):     "hexagonal/main.go",
		".env":                                                "common/env",
		"internal/adapters/cache/cache.go":                    "hexagonal/cache.go",
		"pkg/logger/logger.go":                                "common/logger",
		"pkg/utils/utils.go":                                  "common/utils",
		"pkg/database/database.go":                            "common/database",
		"configs/config.go":                                   "common/config",
	}
	
	return h.renderAndWriteFiles(projectPath, files, templateData)
}

func (h *HexagonalArchitecture) renderAndWriteFiles(projectPath string, files map[string]string, data *TemplateData) error {
	for filePath, templateName := range files {
		content, err := h.templateManager.RenderTemplate(templateName, data)
		if err != nil {
			return fmt.Errorf("failed to render template %s: %w", templateName, err)
		}
		
		fullPath := filepath.Join(projectPath, filePath)
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		
		// Write file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}
	
	return nil
}

// CleanArchitecture implements clean architecture
type CleanArchitecture struct {
	templateManager *TemplateManager
}

func NewCleanArchitecture() *CleanArchitecture {
	tm, err := NewTemplateManager()
	if err != nil {
		panic(fmt.Sprintf("Failed to create template manager: %v", err))
	}
	
	return &CleanArchitecture{
		templateManager: tm,
	}
}

func (c *CleanArchitecture) GetName() string {
	return "clean"
}

func (c *CleanArchitecture) GetStructure() *ProjectStructure {
	structure := NewProjectStructure()
	
	dirs := []string{
		"cmd",
		"app",
		"domain",
		"repository",
		"usecase", 
		"delivery/http",
		"delivery/http/middleware",
		"infrastructure/database",
		"infrastructure/repository",
		"pkg/logger",
		"pkg/utils",
		"configs",
		"docs",
		"scripts",
	}
	
	for _, dir := range dirs {
		structure.AddDirectory(dir)
	}
	
	return structure
}

func (c *CleanArchitecture) GenerateFiles(projectPath string, config *Config) error {
	templateData := NewTemplateData(config)
	
	files := map[string]string{
		fmt.Sprintf("cmd/%s/main.go", config.ProjectName): "clean/main.go",
		".env":                                            "common/env",
		"pkg/logger/logger.go":                            "common/logger",
		"pkg/utils/utils.go":                              "common/utils",
		"pkg/database/database.go":                        "common/database",
		"configs/config.go":                               "common/config",
	}
	
	return c.renderAndWriteFiles(projectPath, files, templateData)
}

func (c *CleanArchitecture) renderAndWriteFiles(projectPath string, files map[string]string, data *TemplateData) error {
	for filePath, templateName := range files {
		content, err := c.templateManager.RenderTemplate(templateName, data)
		if err != nil {
			return fmt.Errorf("failed to render template %s: %w", templateName, err)
		}
		
		fullPath := filepath.Join(projectPath, filePath)
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		
		// Write file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}
	
	return nil
}

// MVCArchitecture implements MVC pattern
type MVCArchitecture struct {
	templateManager *TemplateManager
}

func NewMVCArchitecture() *MVCArchitecture {
	tm, err := NewTemplateManager()
	if err != nil {
		panic(fmt.Sprintf("Failed to create template manager: %v", err))
	}
	
	return &MVCArchitecture{
		templateManager: tm,
	}
}

func (m *MVCArchitecture) GetName() string {
	return "mvc"
}

func (m *MVCArchitecture) GetStructure() *ProjectStructure {
	structure := NewProjectStructure()
	
	dirs := []string{
		"cmd",
		"app",
		"controllers",
		"models",
		"views",
		"middleware",
		"routes",
		"database",
		"migrations",
		"seeders",
		"pkg/logger",
		"pkg/utils",
		"pkg/validators",
		"configs",
		"public/assets",
		"storage/logs",
		"tests",
	}
	
	for _, dir := range dirs {
		structure.AddDirectory(dir)
	}
	
	return structure
}

func (m *MVCArchitecture) GenerateFiles(projectPath string, config *Config) error {
	templateData := NewTemplateData(config)
	
	files := map[string]string{
		fmt.Sprintf("cmd/%s/main.go", config.ProjectName): "mvc/main.go",
		".env":                                            "common/env",
		"pkg/logger/logger.go":                            "common/logger",
		"pkg/utils/utils.go":                              "common/utils",
		"pkg/database/database.go":                        "common/database",
		"configs/config.go":                               "common/config",
	}
	
	return m.renderAndWriteFiles(projectPath, files, templateData)
}

func (m *MVCArchitecture) renderAndWriteFiles(projectPath string, files map[string]string, data *TemplateData) error {
	for filePath, templateName := range files {
		content, err := m.templateManager.RenderTemplate(templateName, data)
		if err != nil {
			return fmt.Errorf("failed to render template %s: %w", templateName, err)
		}
		
		fullPath := filepath.Join(projectPath, filePath)
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		
		// Write file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}
	
	return nil
}

// BasicArchitecture implements basic Go project structure
type BasicArchitecture struct {
	templateManager *TemplateManager
}

func NewBasicArchitecture() *BasicArchitecture {
	tm, err := NewTemplateManager()
	if err != nil {
		panic(fmt.Sprintf("Failed to create template manager: %v", err))
	}
	
	return &BasicArchitecture{
		templateManager: tm,
	}
}

func (b *BasicArchitecture) GetName() string {
	return "basic"
}

func (b *BasicArchitecture) GetStructure() *ProjectStructure {
	structure := NewProjectStructure()
	
	dirs := []string{
		"cmd",
		"internal/app",
		"internal/handlers",
		"internal/services",
		"internal/repository",
		"pkg/logger",
		"pkg/utils",
		"pkg/database",
		"configs",
		"docs",
		"scripts",
		"tests",
	}
	
	for _, dir := range dirs {
		structure.AddDirectory(dir)
	}
	
	return structure
}

func (b *BasicArchitecture) GenerateFiles(projectPath string, config *Config) error {
	templateData := NewTemplateData(config)
	
	files := map[string]string{
		fmt.Sprintf("cmd/%s/main.go", config.ProjectName): "basic/main.go",
		".env":                                            "common/env",
		"pkg/logger/logger.go":                            "common/logger",
		"pkg/utils/utils.go":                              "common/utils",
		"pkg/database/database.go":                        "common/database",
		"configs/config.go":                               "common/config",
	}
	
	return b.renderAndWriteFiles(projectPath, files, templateData)
}

func (b *BasicArchitecture) renderAndWriteFiles(projectPath string, files map[string]string, data *TemplateData) error {
	for filePath, templateName := range files {
		content, err := b.templateManager.RenderTemplate(templateName, data)
		if err != nil {
			return fmt.Errorf("failed to render template %s: %w", templateName, err)
		}
		
		fullPath := filepath.Join(projectPath, filePath)
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		
		// Write file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}
	
	return nil
}
