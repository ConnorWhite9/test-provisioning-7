package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orchctl",
	Short: "orchctl — CLI for the container orchestration system",
	Long: `orchctl is a command-line tool for managing workloads and nodes
in the container orchestration system. Communicates with the control plane
REST API to create, inspect, scale, and delete workloads.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
