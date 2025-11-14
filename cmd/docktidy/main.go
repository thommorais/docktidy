// Package main provides the CLI entry point for docktidy.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thommorais/docktidy/internal/adapters/docker"
	"github.com/thommorais/docktidy/internal/adapters/tui"
	"github.com/thommorais/docktidy/internal/domain"
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
		status, usage := dockerStatusAndUsage()
		app := tui.New(
			tui.WithDockerStatus(status),
			tui.WithDiskUsage(usage),
		)
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

func dockerStatusAndUsage() (tui.StatusMessage, domain.DiskUsage) {
	txt := text.Default()
	status := tui.StatusMessage{
		Message: txt.Get(text.KeyDockerStatusUnknown),
		Level:   tui.StatusLevelUnknown,
	}
	var usage domain.DiskUsage

	svc, err := docker.NewService()
	if err != nil {
		return degradedStatus(status, txt, err), usage
	}

	ctx := context.Background()
	if err := svc.IsHealthy(ctx); err != nil {
		return degradedStatus(status, txt, err), usage
	}

	data, err := svc.DiskUsage(ctx)
	if err != nil {
		return degradedStatus(status, txt, err), usage
	}

	status.Message = txt.Get(text.KeyDockerStatusHealthy)
	status.Level = tui.StatusLevelHealthy
	return status, data
}

func degradedStatus(base tui.StatusMessage, txt *text.Text, err error) tui.StatusMessage {
	base.Message = fmt.Sprintf("%s (%v)", txt.Get(text.KeyDockerStatusDegraded), err)
	base.Level = tui.StatusLevelDegraded
	return base
}
