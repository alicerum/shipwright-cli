package cmd

import (
	"github.com/shipwright-io/cli/pkg/shp/cmd/build"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var rootCmd = &cobra.Command{
	Use:   "shp [command] [resource] [flags]",
	Short: "Command-line client for Shipwright's Build API.",
}

// NewCmdSHP create a new SHP root command, linking together all sub-commands organized by groups.
func NewCmdSHP(ioStreams genericclioptions.IOStreams) *cobra.Command {
	opts := runner.NewOptions()
	// wiring up root command flags with options instance
	opts.AddFlags(rootCmd.Flags())

	rootCmd.AddCommand(build.Command())

	return rootCmd
}
