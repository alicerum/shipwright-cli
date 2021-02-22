package build

import (
	"fmt"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type ListSubCommand struct {
	cmd *cobra.Command
}

func (sc *ListSubCommand) Cmd() *cobra.Command {
	return sc.cmd
}

func (sc *ListSubCommand) Complete(params params.Params, args []string) error {
	return nil
}

func (sc *ListSubCommand) Validate() error {
	return nil
}

func (sc *ListSubCommand) Run(params params.Params) error {
	fmt.Printf("ns is %v\n", params.Namespace())
	_, err := params.Client()
	if err != nil {
		return err
	}

	return nil
}

func listSubCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "list [flags]",
		Short: "List Builds",
	}

	return &ListSubCommand{
		cmd: c,
	}
}
