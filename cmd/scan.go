package cmd

import (
	"fmt"
	"os"

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
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
