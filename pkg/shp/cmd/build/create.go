package build

import (
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type CreateSubCommand struct {
	cmd *cobra.Command
}

func (sc *CreateSubCommand) Cmd() *cobra.Command {
	return sc.cmd
}

func (sc *CreateSubCommand) Complete(params params.Params, args []string) error {
	return nil
}

func (sc *CreateSubCommand) Validate() error {
	return nil
}

func (sc *CreateSubCommand) Run(params params.Params) error {
	return nil
}

func createSubCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "create [flags]",
		Short: "Create Build",
	}

	return &CreateSubCommand{
		cmd: c,
	}
}
