package build

import (
	"errors"
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/cmd/buildrun"
	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type RunSubCommand struct {
	cmd *cobra.Command

	buildName string
}

func (sc *RunSubCommand) Cmd() *cobra.Command {
	return sc.cmd
}

func runSubCmd() runner.SubCommand {
	cmd := &cobra.Command{
		Use:   "run [name]",
		Short: "Start a build specified by 'name'",
	}

	return &RunSubCommand{
		cmd: cmd,
	}
}

func (sc *RunSubCommand) Complete(params params.Params, args []string) error {
	if len(args) < 1 {
		return errors.New("'name' argument is empty")
	}

	sc.buildName = args[0]

	return nil
}

func (sc *RunSubCommand) Validate() error {
	return nil
}

func (sc *RunSubCommand) Run(params params.Params) error {
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
