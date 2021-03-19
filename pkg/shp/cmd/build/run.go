package build

import (
	"errors"
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/spf13/cobra"

	"github.com/shipwright-io/cli/pkg/shp/cmd/buildrun"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
)

type RunCommand struct {
	cmd *cobra.Command

	buildName string
}

func (c *RunCommand) Cmd() *cobra.Command {
	return c.cmd
}

func runCmd() runner.SubCommand {
	cmd := &cobra.Command{
		Use:   "run [name]",
		Short: "Start a build specified by 'name'",
	}

	return &RunCommand{
		cmd: cmd,
	}
}

func (c *RunCommand) Complete(params *params.Params, args []string) error {
	if len(args) < 1 {
		return errors.New("'name' argument is empty")
	}

	c.buildName = args[0]

	return nil
}

func (c *RunCommand) Validate() error {
	return nil
}

func (sc *RunCommand) Run(params *params.Params) error {
	br := GetBuildResource(params)
	brr := buildrun.GetBuildRunResource(params)

	var build buildv1alpha1.Build
	if err := br.Get(sc.cmd.Context(), sc.buildName, &build); err != nil {
		return err
	}

	buildRun := buildrun.NewBuildRun(&build, build.Name+"-run")

	if err := brr.Create(sc.cmd.Context(), build.Name, buildRun); err != nil {
		return err
	}

	fmt.Printf("BuildRun created \"%v\"\n", buildRun.Name)
	return nil
}
