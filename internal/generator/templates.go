package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

// TemplateManager manages template loading and rendering
type TemplateManager struct {
	templates map[string]*template.Template
}

// NewTemplateManager creates a new template manager
func NewTemplateManager() (*TemplateManager, error) {
	tm := &TemplateManager{
		templates: make(map[string]*template.Template),
	}

	if err := tm.loadTemplates(); err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return tm, nil
}

// loadTemplates loads all templates from embedded filesystem
func (tm *TemplateManager) loadTemplates() error {
	return filepath.WalkDir("templates", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		content, err := templatesFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		// Create template name from path (remove templates/ prefix and .tmpl suffix)
		templateName := strings.TrimPrefix(path, "templates/")
		templateName = strings.TrimSuffix(templateName, ".tmpl")

		tmpl, err := template.New(templateName).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", path, err)
		}

		tm.templates[templateName] = tmpl
		return nil
	})
}

// RenderTemplate renders a template with given data
func (tm *TemplateManager) RenderTemplate(templateName string, data interface{}) (string, error) {
	tmpl, exists := tm.templates[templateName]
	if !exists {
		return "", fmt.Errorf("template %s not found", templateName)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return buf.String(), nil
}

// GetTemplate returns a template by name
func (tm *TemplateManager) GetTemplate(templateName string) (*template.Template, error) {
	tmpl, exists := tm.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template %s not found", templateName)
	}
	return tmpl, nil
}

// ListTemplates returns all available template names
func (tm *TemplateManager) ListTemplates() []string {
	var names []string
	for name := range tm.templates {
		names = append(names, name)
	}
	return names
}
