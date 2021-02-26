package build

import (
	"errors"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/util"
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
	client, err := params.Client()
	if err != nil {
		return err
	}

	var b buildv1alpha1.Build
	res := client.Resource(BuildGVR).Namespace(params.Namespace())

	err = util.GetObject(res, sc.name, BuildGVR, &b)
	if err != nil {
		return err
	}

	return util.DeleteObject(res, sc.name, BuildGVR)
}
