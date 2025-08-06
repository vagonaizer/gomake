package cli

import (
	"fmt"

	"github.com/gomake/internal/generator"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  "Print version information for gomake",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generator.GetVersionInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
