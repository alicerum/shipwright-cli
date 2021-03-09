package build

import (
	"os"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
)

// Command returns Build subcommand of Shipwright CLI
// for interaction with shipwright builds
func Command(p *params.Params) *cobra.Command {
	command := &cobra.Command{
		Use:     "build",
		Aliases: []string{"bd"},
		Short:   "Manage Builds",
		Annotations: map[string]string{
			"commandType": "main",
		},
	}

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	command.AddCommand(
		runner.NewRunner(p, streams, createCmd()).Cmd(),
		runner.NewRunner(p, streams, listCmd()).Cmd(),
		runner.NewRunner(p, streams, deleteCmd()).Cmd(),
		runner.NewRunner(p, streams, runCmd()).Cmd(),
	)
	return command
}
