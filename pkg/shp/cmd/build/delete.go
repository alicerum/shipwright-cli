package build

import (
	"errors"
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/cmd/buildrun"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteCommand struct {
	name string

	cmd        *cobra.Command
	deleteRuns bool
}

func deleteCmd() runner.SubCommand {
	deleteCmd := &DeleteCommand{
		cmd: &cobra.Command{
			Use:   "delete [flags] [name]",
			Short: "Delete Build",
		},
	}

	deleteCmd.cmd.Flags().BoolVarP(&deleteCmd.deleteRuns, "delete-runs", "r", false, "Also delete all the buildruns")

	return deleteCmd
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
	brr := buildrun.GetBuildRunResource(params)

	if err := br.Get(c.cmd.Context(), c.name, &b); err != nil {
		return err
	}

	if err := br.Delete(c.cmd.Context(), c.name); err != nil {
		return err
	}

	if c.deleteRuns {
		label := fmt.Sprintf("build.build.dev/name=%v", c.name)
		var brList buildv1alpha1.BuildRunList
		err := brr.ListWithOptions(c.cmd.Context(), &brList, v1.ListOptions{
			LabelSelector: label,
		})

		if err != nil {
			return err
		}

		for _, buildrun := range brList.Items {
			brr.Delete(c.cmd.Context(), buildrun.Name)
		}
	}

	return nil
}
