package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "infra-lens",
	Short: "Audit your Terraform infrastructure",
	Long:  `infra-lens scans Terraform files and surfaces security issues, cost inefficiencies, and technical debt.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
}
