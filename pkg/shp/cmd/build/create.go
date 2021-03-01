package build

import (
	"errors"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

// CreateCommand contains data input from user
type CreateCommand struct {
	url      string
	strategy string
	name     string

	image string

	build *buildv1alpha1.Build

	cmd *cobra.Command
}

func createCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "create [flags] name strategy url",
		Short: "Create Build",
	}

	c.Flags().StringP("output-image", "i", "", "Output image created by build")

	return &CreateCommand{
		cmd: c,
	}
}

func (c *CreateCommand) Cmd() *cobra.Command {
	return c.cmd
}

func (c *CreateCommand) Complete(params *params.Params, args []string) error {

	if len(args) < 3 {
		return errors.New("Not enough arguments for Build create")
	}

	c.name = args[0]
	c.strategy = args[1]
	c.url = args[2]

	return nil
}

func (c *CreateCommand) initializeBuild() {
	strategyKind := buildv1alpha1.ClusterBuildStrategyKind

	c.build = &buildv1alpha1.Build{
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

	if c.image != "" {
		c.build.Spec.Output = buildv1alpha1.Image{
			ImageURL: c.image,
		}
	}
}

func (c *CreateCommand) Validate() error {
	if c.strategy != "buildah" {
		return errors.New("Incorrect strategy, must be 'buildah'")
	}

	return nil
}

func (c *CreateCommand) Run(params params.Params) error {
	c.initializeBuild()

	return buildResource.Create(c.name, c.build)
}
