package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gomake/internal/generator"
	"github.com/gomake/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	// Global flags
	architecture string
	autoYes      bool
	targetDir    string
	withDocker   bool
	withMakefile bool
	withGit      bool
	interactive  bool
	license      string
	verbose      bool

	// Logger instance
	log *logger.Logger

	// Available architectures
	availableArchs = []string{"hexagonal", "clean", "mvc", "basic"}
)

var rootCmd = &cobra.Command{
	Use:   "gomake",
	Short: "A CLI tool for generating Go project structures",
	Long: color.CyanString(`
╔═══════════════════════════════════════╗
║              GOMAKE v1.0              ║
║   Go Project Structure Generator      ║
╚═══════════════════════════════════════╝
`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log = logger.New(verbose)
	},
}

var projectCmd = &cobra.Command{
	Use:   "project [project-name]",
	Short: "Generate a new Go project",
	Long:  "Generate a new Go project with the specified architecture and options",
	Args:  cobra.ExactArgs(1),
	RunE:  runProjectCommand,
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(projectCmd)

	// Project command flags
	projectCmd.Flags().StringVarP(&architecture, "arch", "a", "basic",
		fmt.Sprintf("Architecture type (%v)", availableArchs))
	projectCmd.Flags().BoolVarP(&autoYes, "yes", "y", false,
		"Automatic confirmation without prompts")
	projectCmd.Flags().StringVarP(&targetDir, "dir", "d", ".",
		"Target directory for project creation")
	projectCmd.Flags().BoolVar(&withDocker, "with-docker", false,
		"Add Dockerfile and docker-compose.yml")
	projectCmd.Flags().BoolVar(&withMakefile, "with-makefile", false,
		"Add Makefile with common targets (always included by default)")
	projectCmd.Flags().BoolVar(&withGit, "with-git", false,
		"Initialize git repository")
	projectCmd.Flags().BoolVarP(&interactive, "interactive", "i", false,
		"Interactive mode with step-by-step wizard")
	projectCmd.Flags().StringVarP(&license, "license", "l", "MIT",
		"License type (MIT, Apache, BSD, GPL)")

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"Verbose output")
}

func runProjectCommand(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	log.Info("Starting project generation", "project", projectName)

	// Validate inputs
	if err := validateInputs(projectName); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Interactive mode
	if interactive {
		if err := runInteractiveMode(&projectName); err != nil {
			return fmt.Errorf("interactive mode failed: %w", err)
		}
	}

	// Create generator config
	config := &generator.Config{
		ProjectName:  projectName,
		Architecture: architecture,
		TargetDir:    targetDir,
		WithDocker:   withDocker,
		WithMakefile: withMakefile,
		WithGit:      withGit,
		License:      license,
		AutoYes:      autoYes,
	}

	// Create generator
	gen, err := generator.New(config, log)
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}

	// Generate project
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	// Success message
	color.Green("\n✅ Project '%s' generated successfully!", projectName)
	color.Cyan("📁 Location: %s/%s", targetDir, projectName)
	color.Yellow("🚀 Next steps:")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Printf("   go mod tidy\n")
	fmt.Printf("   make help\n")
	fmt.Printf("   make run\n")

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

func boolToString(b bool) string {
	if b {
		return color.GreenString("Yes")
	}
	return color.RedString("No")
}
