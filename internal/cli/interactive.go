package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func runInteractiveMode(projectName *string) error {
	reader := bufio.NewReader(os.Stdin)

	color.Cyan("\nInteractive Project Setup")
	color.Cyan("═══════════════════════════\n")

	// Project name confirmation
	if *projectName == "" {
		fmt.Print("// Enter project name: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		*projectName = strings.TrimSpace(name)
	} else {
		fmt.Printf(" // Project name: %s\n", color.GreenString(*projectName))
		if !askConfirmation(reader, "Continue with this name?") {
			fmt.Print("📝 Enter new project name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			*projectName = strings.TrimSpace(name)
		}
	}

	// Architecture selection
	color.Yellow("\n // Select Architecture:")
	for i, arch := range availableArchs {
		fmt.Printf("   %d) %s\n", i+1, arch)
	}

	fmt.Print("Enter choice (1-4): ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	choiceNum, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil || choiceNum < 1 || choiceNum > len(availableArchs) {
		color.Red("Invalid choice, using 'basic'")
		architecture = "basic"
	} else {
		architecture = availableArchs[choiceNum-1]
	}

	// Additional options
	color.Yellow("\n// Additional Options:")
	withDocker = askConfirmation(reader, "Add Docker support?")
	withMakefile = askConfirmation(reader, "Add Makefile?")
	withGit = askConfirmation(reader, "Initialize Git repository?")

	// License selection
	color.Yellow("\n// License Selection:")
	licenses := []string{"MIT", "Apache", "BSD", "GPL", "None"}
	for i, lic := range licenses {
		fmt.Printf("   %d) %s\n", i+1, lic)
	}

	fmt.Print("Enter choice (1-5): ")
	licChoice, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	licChoiceNum, err := strconv.Atoi(strings.TrimSpace(licChoice))
	if err != nil || licChoiceNum < 1 || licChoiceNum > len(licenses) {
		license = "MIT"
	} else {
		license = licenses[licChoiceNum-1]
	}

	// Summary
	color.Cyan("\n// Project Summary:")
	fmt.Printf("   Name: %s\n", color.GreenString(*projectName))
	fmt.Printf("   Architecture: %s\n", color.GreenString(architecture))
	fmt.Printf("   Docker: %s\n", boolToString(withDocker))
	fmt.Printf("   Makefile: %s\n", boolToString(withMakefile))
	fmt.Printf("   Git: %s\n", boolToString(withGit))
	fmt.Printf("   License: %s\n", color.GreenString(license))

	return nil
}

func askConfirmation(reader *bufio.Reader, question string) bool {
	fmt.Printf("%s (y/N): ", question)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}
