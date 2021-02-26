package runner

import (
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Runner execute the sub-command lifecycle, wrapper around sub-commands.
type Runner struct {
	opts      *Options                    // global options
	ioStreams genericclioptions.IOStreams // input, output and error io streams
	subCmd    SubCommand                  // sub-command instance
}

// Cmd is a wrapper around sub-command's Cobra, it wires up global flags and set a single RunE
// executor to self.
func (r *Runner) Cmd() *cobra.Command {
	cmd := r.subCmd.Cmd()
	cmd.RunE = r.RunE
	r.opts.AddFlags(cmd.PersistentFlags())
	return cmd
}

func (r *Runner) createParams() (params.Params, error) {
	namespace := *r.opts.configFlags.Namespace
	context := *r.opts.configFlags.Context
	configPath := *r.opts.configFlags.KubeConfig

	return params.NewParams(configPath, namespace, context)
}

// RunE cobra.Command's RunE implementation focusing on sub-commands lifecycle. To achieve it, a
// dynamic client and configured namespace are informed.
func (r *Runner) RunE(cmd *cobra.Command, args []string) error {
	params, err := r.createParams()
	if err != nil {
		return err
	}

	if err = r.subCmd.Complete(params, args); err != nil {
		return err
	}
	if err = r.subCmd.Validate(); err != nil {
		return err
	}
	return r.subCmd.Run(params)
}

// NewRunner instantiate a Runner.
func NewRunner(opts *Options, ioStreams genericclioptions.IOStreams, subCmd SubCommand) *Runner {
	return &Runner{opts: opts, ioStreams: ioStreams, subCmd: subCmd}
}