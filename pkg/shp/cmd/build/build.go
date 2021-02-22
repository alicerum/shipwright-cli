package build

import (
	"os"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:     "build",
		Aliases: []string{"bd", "build"},
		Short:   "Manage Builds",
		Annotations: map[string]string{
			"commandType": "main",
		},
	}

	o := runner.NewOptions()
	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	command.AddCommand(
		runner.NewRunner(o, streams, createSubCmd()).Cmd(),
		runner.NewRunner(o, streams, listSubCmd()).Cmd(),
	)
	return command
}
