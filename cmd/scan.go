package cmd

import (
	"fmt"
	"os"

	"github.com/Baixerastheo/infra-lens/internal/parser"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a Terraform directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error: path '%s' does not exist\n", path)
			os.Exit(1)
		}

		fmt.Printf("Scanning %s...\n", path)

		resources, err := parser.Parse(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Found %d ressources\n", len(resources))
		for _, r := range resources {
			fmt.Printf(" → %s.%s (%s)\n", r.Type, r.Name, r.File)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
