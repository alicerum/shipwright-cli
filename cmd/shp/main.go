package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	kubeContext string
	dryRun      bool
	kubeconfig  string
	namespace   string
)

var rootCmd = &cobra.Command{
	Use:   "shp [command]",
	Short: "Command-line client for Shipwright Build Operator.",
}

// init setup flags on the root command.
func init() {
	flags := rootCmd.PersistentFlags()

	flags.StringVar(
		&kubeContext,
		"context",
		"",
		"alternative Kubernetes context, when empty it use default context in configuration",
	)
	flags.StringVar(
		&kubeconfig,
		"kubeconfig",
		defaultKubeconfigPath(),
		"path to kubeconfig file",
	)
	flags.StringVar(
		&namespace,
		"namespace",
		"",
		"alternative namespace, when empty it use namespace configured for context",
	)
	flags.BoolVar(
		&dryRun,
		"dry-run",
		false,
		"avoid any updates on the cluster resources",
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
