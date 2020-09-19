package cmd

import (
	"github.com/otaviof/shp/pkg/shp/buildrun"
	"github.com/otaviof/shp/pkg/shp/initialize"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/templates"
)

var rootCmd = &cobra.Command{
	Use:   "shp [command] [resource] [flags]",
	Short: "Command-line client for Shipwright's Build API.",
}

// NewCmdSHP create a new SHP root command, linking together all sub-commands organized by groups.
func NewCmdSHP(ioStreams genericclioptions.IOStreams) *cobra.Command {
	opts := NewOptions()
	// wiring up root command flags with options instance
	opts.AddFlags(rootCmd.Flags())

	cg := templates.CommandGroups{{
		Message: "Initialize repository commands:",
		Commands: []*cobra.Command{
			NewRunner(opts, ioStreams, initialize.NewInitialize()).Cmd(),
		},
	}, {
		Message: "Manage BuildRun Resources:",
		Commands: []*cobra.Command{
			NewRunner(opts, ioStreams, buildrun.NewCreateBuildRun()).Cmd(),
			NewRunner(opts, ioStreams, buildrun.NewDeleteBuildRun()).Cmd(),
		},
	}, {
		Message: "Manage Build Resources:",
		Commands: []*cobra.Command{
			NewRunner(opts, ioStreams, buildrun.NewRunBuild()).Cmd(),
		},
	}}
	// linking command groups with root command
	cg.Add(rootCmd)

	return rootCmd
}
