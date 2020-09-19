package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
)

type Runner struct {
	opts      *Options
	ioStreams genericclioptions.IOStreams
	subCmd    SubCommand
}

// Cmd is a wrapper around sub-command's Cobra, it wires up global flags and set a single RunE
// executor to self.
func (r *Runner) Cmd() *cobra.Command {
	cmd := r.subCmd.Cmd()
	cmd.RunE = r.RunE
	r.opts.AddFlags(cmd.PersistentFlags())
	return cmd
}

func (r *Runner) dynamicClientNamespace() (dynamic.Interface, string, error) {
	f := r.opts.Factory()
	configLoader := f.ToRawKubeConfigLoader()

	namespace := *r.opts.configFlags.Namespace
	if namespace == "" {
		var err error
		ns, _, err := configLoader.Namespace()
		if err != nil {
			return nil, "", err
		}
		namespace = ns
	}

	restConfig, err := configLoader.ClientConfig()
	if err != nil {
		return nil, "", err
	}
	client, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, "", err
	}
	return client, namespace, nil
}

func (r *Runner) RunE(cmd *cobra.Command, args []string) error {
	client, ns, err := r.dynamicClientNamespace()
	if err != nil {
		return err
	}

	if err = r.subCmd.Complete(client, ns, args); err != nil {
		return err
	}
	if err = r.subCmd.Validate(); err != nil {
		return err
	}
	return r.subCmd.Run(client, ns)
}

func NewRunner(opts *Options, ioStreams genericclioptions.IOStreams, subCmd SubCommand) *Runner {
	return &Runner{opts: opts, ioStreams: ioStreams, subCmd: subCmd}
}
