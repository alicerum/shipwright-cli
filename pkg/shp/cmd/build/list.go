package build

import (
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/util"
	"github.com/spf13/cobra"
)

type ListSubCommand struct {
	cmd *cobra.Command
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
	client, err := params.Client()
	if err != nil {
		return err
	}

	res := client.Resource(BuildGVR).Namespace(params.Namespace())
	var buildList buildv1alpha1.BuildList
	if err = util.ListObject(res, &buildList); err != nil {
		return err
	}

	for _, b := range buildList.Items {
		fmt.Printf("%v\t%v\n", b.Name, b.Status.Message)
	}

	return nil
}
