package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gomake/internal/generator"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage gomake configuration",
	Long:  "Manage gomake configuration files and templates",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize default configuration file",
	Long:  "Create a default .gomake.yml configuration file in the current directory",
	RunE:  runConfigInit,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  "Display the current gomake configuration",
	RunE:  runConfigShow,
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	rootCmd.AddCommand(configCmd)
}

func runConfigInit(cmd *cobra.Command, args []string) error {
	configPath := ".gomake.yml"

	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		color.Yellow("⚠️  Configuration file already exists: %s", configPath)
		fmt.Print("Do you want to overwrite it? (y/N): ")

		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "yes" {
			color.Blue("ℹ️  Configuration initialization cancelled")
			return nil
		}
	}

	// Create default config
	if err := generator.SaveDefaultConfig(configPath); err != nil {
		return fmt.Errorf("failed to create configuration file: %w", err)
	}

	color.Green("✅ Configuration file created: %s", configPath)
	color.Cyan("📝 You can now customize your templates and defaults")

	return nil
}

func runConfigShow(cmd *cobra.Command, args []string) error {
	config, err := generator.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	color.Cyan("📋 Current Configuration:")
	fmt.Printf("\n")

	// Show defaults
	color.Yellow("🔧 Defaults:")
	fmt.Printf("  Architecture: %s\n", config.Defaults.Architecture)
	fmt.Printf("  License: %s\n", config.Defaults.License)
	fmt.Printf("  Docker: %v\n", config.Defaults.WithDocker)
	fmt.Printf("  Makefile: %v\n", config.Defaults.WithMakefile)
	fmt.Printf("  Git: %v\n", config.Defaults.WithGit)

	// Show custom templates
	if len(config.Templates) > 0 {
		fmt.Printf("\n")
		color.Yellow("📦 Custom Templates:")
		for _, template := range config.Templates {
			fmt.Printf("  • %s: %s\n", template.Name, template.Description)
		}
	} else {
		fmt.Printf("\n")
		color.Blue("ℹ️  No custom templates configured")
	}

	// Show config file locations
	fmt.Printf("\n")
	color.Yellow("📁 Config file locations (in order of precedence):")
	configPaths := []string{
		".gomake.yml",
		".gomake.yaml",
		filepath.Join(os.Getenv("HOME"), ".gomake.yml"),
		filepath.Join(os.Getenv("HOME"), ".config", "gomake", "config.yml"),
	}

	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			color.Green("  ✓ %s (found)", path)
		} else {
			fmt.Printf("  • %s\n", path)
		}
	}

	return nil
}
