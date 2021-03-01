package build

import (
	"errors"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type DeleteSubCommand struct {
	name string

	cmd *cobra.Command
}

func deleteSubCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "delete [flags] name",
		Short: "Delete Build",
	}

	return &DeleteSubCommand{
		cmd: c,
	}
}

func (sc *DeleteSubCommand) Cmd() *cobra.Command {
	return sc.cmd
}

func (sc *DeleteSubCommand) Complete(params params.Params, args []string) error {
	if len(args) < 1 {
		return errors.New("Missing 'name' argument")
	}

	sc.name = args[0]

	return nil
}

func (sc *DeleteSubCommand) Validate() error {
	return nil
}

func (sc *DeleteSubCommand) Run(params params.Params) error {
	var b buildv1alpha1.Build

	if err := buildResource.Get(sc.name, &b); err != nil {
		return err
	}

	return buildResource.Delete(sc.name)
}
