package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thommorais/docktidy/internal/adapters/tui"
)

var (
	// Version information (set by goreleaser)
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "docktidy",
	Short: "A TUI tool for safely cleaning up Docker resources",
	Long: `docktidy helps you identify and safely remove unused Docker resources
including containers, images, volumes, and networks.

It tracks resource usage history and provides intelligent suggestions
for what can be safely pruned to reclaim disk space.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Start the TUI
		app := tui.New()
		return app.Run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("docktidy %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built:  %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Version = version
	// Add global flags here in the future
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docktidy.yaml)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
