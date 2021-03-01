package build

import (
	"os"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/resource"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	buildResource *resource.ShpResource
)

func Command(p params.Params) *cobra.Command {
	buildResource = resource.NewShpResource(
		p,
		buildv1alpha1.SchemeGroupVersion,
		"Build",
		"builds",
	)

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
		runner.NewRunner(p, streams, createSubCmd()).Cmd(),
		runner.NewRunner(p, streams, listSubCmd()).Cmd(),
		runner.NewRunner(p, streams, deleteSubCmd()).Cmd(),
	)
	return command
}
