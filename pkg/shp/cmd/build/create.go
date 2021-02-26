package build

import (
	"errors"
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/util"
	"github.com/spf13/cobra"
)

type CreateSubCommand struct {
	url      string
	strategy string
	name     string

	image string

	build *buildv1alpha1.Build

	cmd *cobra.Command
}

func createSubCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "create [flags] name strategy url",
		Short: "Create Build",
	}

	c.Flags().StringP("image", "i", "", "Output image created by build")

	return &CreateSubCommand{
		cmd: c,
	}
}

func (sc *CreateSubCommand) Cmd() *cobra.Command {
	return sc.cmd
}

func (sc *CreateSubCommand) Complete(params params.Params, args []string) error {

	if len(args) < 3 {
		return errors.New("Not enough arguments for Build create")
	}

	sc.name = args[0]
	sc.strategy = args[1]
	sc.url = args[2]

	sc.image = sc.cmd.Flag("image").Value.String()

	return nil
}

func (sc *CreateSubCommand) initializeBuild() {
	strategyKind := buildv1alpha1.ClusterBuildStrategyKind

	sc.build = &buildv1alpha1.Build{
		ObjectMeta: metav1.ObjectMeta{
			Name: sc.name,
		},
		Spec: buildv1alpha1.BuildSpec{
			StrategyRef: &buildv1alpha1.StrategyRef{
				Name: sc.strategy,
				Kind: &strategyKind,
			},
			Source: buildv1alpha1.GitSource{
				URL: sc.url,
			},
		},
	}

	if sc.image != "" {
		sc.build.Spec.Output = buildv1alpha1.Image{
			ImageURL: sc.image,
		}
	}
}

func (sc *CreateSubCommand) Validate() error {
	if sc.strategy != "buildah" {
		return errors.New("Incorrect strategy, can be 'buildah'")
	}

	return nil
}

func (sc *CreateSubCommand) Run(params params.Params) error {
	fmt.Println("Url is " + sc.url)

	sc.initializeBuild()

	client, err := params.Client()
	if err != nil {
		return nil
	}

	res := client.Resource(BuildGVR).Namespace(params.Namespace())
	return util.CreateObject(res, sc.name, BuildGVK, sc.build)
}
