package build

import (
	"fmt"
	"os"
	"text/tabwriter"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"

	"github.com/shipwright-io/cli/pkg/shp/cmd/runner"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
)

type ListSubCommand struct {
	cmd *cobra.Command

	noHeader *bool
}

var (
	writer *tabwriter.Writer

	columnNames    string
	columnTemplate string
)

func init() {
	writer = tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)

	columnNames = "NAME\tOUTPUT\tSTATUS"
	columnTemplate = "%s\t%s\t%s\n"
}

func listSubCmd() runner.SubCommand {
	c := &cobra.Command{
		Use:   "list [flags]",
		Short: "List Builds",
	}

	noHeader := c.Flags().Bool("no-header", false, "Do not show columns header in list output")

	return &ListSubCommand{
		cmd:      c,
		noHeader: noHeader,
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
	var buildList buildv1alpha1.BuildList
	if err := buildResource.List(&buildList); err != nil {
		return err
	}

	if !*sc.noHeader {
		fmt.Fprintln(writer, columnNames)
	}

	for _, b := range buildList.Items {
		fmt.Fprintf(writer, columnTemplate, b.Name, b.Spec.Output.ImageURL, b.Status.Message)
	}

	writer.Flush()

	return nil
}
