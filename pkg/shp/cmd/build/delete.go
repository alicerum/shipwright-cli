package build

import (
	"errors"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type DeleteCommand struct {
	name string

	cmd *cobra.Command
}

func deleteCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "delete [flags] [name]",
		Short: "Delete Build",
	}

	return &DeleteCommand{
		cmd: c,
	}
}

func (s *DeleteCommand) Cmd() *cobra.Command {
	return s.cmd
}

func (s *DeleteCommand) Complete(params *params.Params, args []string) error {
	if len(args) < 1 {
		return errors.New("Missing 'name' argument")
	}

	s.name = args[0]

	return nil
}

func (c *DeleteCommand) Validate() error {
	return nil
}

func (c *DeleteCommand) Run(params *params.Params) error {
	var b buildv1alpha1.Build

	br := GetBuildResource(params)

	if err := br.Get(c.cmd.Context(), c.name, &b); err != nil {
		return err
	}

	return br.Delete(c.cmd.Context(), c.name)
}
