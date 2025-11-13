// Package main provides the CLI entry point for docktidy.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thommorais/docktidy/internal/adapters/docker"
	"github.com/thommorais/docktidy/internal/adapters/tui"
	"github.com/thommorais/docktidy/pkg/text"
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
	RunE: func(_ *cobra.Command, _ []string) error {
		app := tui.New(tui.WithDockerStatus(dockerHealthStatus()))
		return app.Run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(_ *cobra.Command, _ []string) {
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

func dockerHealthStatus() tui.StatusMessage {
	txt := text.Default()
	status := tui.StatusMessage{
		Message: txt.Get(text.KeyDockerStatusUnknown),
		Level:   tui.StatusLevelUnknown,
	}

	svc, err := docker.NewService()
	if err != nil {
		status.Message = fmt.Sprintf("%s (%v)", txt.Get(text.KeyDockerStatusDegraded), err)
		status.Level = tui.StatusLevelDegraded
		return status
	}

	if err := svc.IsHealthy(context.Background()); err != nil {
		status.Message = fmt.Sprintf("%s (%v)", txt.Get(text.KeyDockerStatusDegraded), err)
		status.Level = tui.StatusLevelDegraded
		return status
	}

	status.Message = txt.Get(text.KeyDockerStatusHealthy)
	status.Level = tui.StatusLevelHealthy
	return status
}
