package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	// Valid Go module name pattern
	moduleNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9._/\-]*[a-zA-Z0-9]$`)
	
	// Reserved Go keywords
	goKeywords = map[string]bool{
		"break": true, "case": true, "chan": true, "const": true, "continue": true,
		"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
		"func": true, "go": true, "goto": true, "if": true, "import": true,
		"interface": true, "map": true, "package": true, "range": true, "return": true,
		"select": true, "struct": true, "switch": true, "type": true, "var": true,
	}
)

func validateInputs(projectName string) error {
	// Validate project name
	if err := validateProjectName(projectName); err != nil {
		return err
	}

	// Validate architecture
	if err := validateArchitecture(); err != nil {
		return err
	}

	// Validate target directory
	if err := validateTargetDirectory(); err != nil {
		return err
	}

	// Check if project already exists
	if err := checkProjectExists(projectName); err != nil {
		return err
	}

	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if len(name) < 2 {
		return fmt.Errorf("project name must be at least 2 characters long")
	}

	if len(name) > 100 {
		return fmt.Errorf("project name must be less than 100 characters")
	}

	// Check for Go keywords
	if goKeywords[strings.ToLower(name)] {
		return fmt.Errorf("project name cannot be a Go keyword: %s", name)
	}

	// Check valid characters
	if !moduleNameRegex.MatchString(name) {
		return fmt.Errorf("invalid project name: %s. Must contain only letters, numbers, dots, underscores, and hyphens", name)
	}

	// Check for invalid patterns
	if strings.HasPrefix(name, ".") || strings.HasSuffix(name, ".") {
		return fmt.Errorf("project name cannot start or end with a dot")
	}

	if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
		return fmt.Errorf("project name cannot start or end with a hyphen")
	}

	return nil
}

func validateArchitecture() error {
	for _, arch := range availableArchs {
		if arch == architecture {
			return nil
		}
	}
	return fmt.Errorf("invalid architecture: %s. Available: %v", architecture, availableArchs)
}

func validateTargetDirectory() error {
	if targetDir == "" {
		targetDir = "."
		return nil
	}

	// Check if directory exists
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return fmt.Errorf("target directory does not exist: %s", targetDir)
	}

	// Check if directory is writable
	testFile := filepath.Join(targetDir, ".gomake_test")
	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("target directory is not writable: %s", targetDir)
	}
	file.Close()
	os.Remove(testFile)

	return nil
}

func checkProjectExists(projectName string) error {
	projectPath := filepath.Join(targetDir, projectName)
	
	if _, err := os.Stat(projectPath); err == nil {
		if autoYes {
			color.Yellow("⚠️  Project directory already exists: %s", projectPath)
			return nil
		}

		color.Yellow("⚠️  Project directory already exists: %s", projectPath)
		fmt.Print("Do you want to overwrite it? (y/N): ")
		
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		
		if response != "y" && response != "yes" {
			return fmt.Errorf("project creation cancelled")
		}
		
		color.Yellow("🗑️  Removing existing directory...")
		if err := os.RemoveAll(projectPath); err != nil {
			return fmt.Errorf("failed to remove existing directory: %w", err)
		}
	}

	return nil
}
