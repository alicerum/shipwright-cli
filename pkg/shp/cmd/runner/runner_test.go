package runner

import (
	"os"
	"testing"

	"github.com/onsi/gomega"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type mockedSubCommand struct{}

var testCmd = &cobra.Command{}

func (m *mockedSubCommand) Cmd() *cobra.Command {
	return testCmd
}

func (m *mockedSubCommand) Complete(p params.Params, args []string) error {
	return nil
}

func (m *mockedSubCommand) Validate() error {
	return nil
}

func (m *mockedSubCommand) Run(p params.Params) error {
	return nil
}

func TestCMD_Runner(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	p := params.NewParams()
	// testNs := "test"
	//p.configFlags.Namespace = &testNs

	genericOpts := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	r := NewRunner(p, genericOpts, &mockedSubCommand{})

	t.Run("cmd", func(t *testing.T) {
		cmd := r.Cmd()

		g.Expect(cmd.RunE).ToNot(gomega.BeNil())
	})

	t.Run("dynamicClientNamespace", func(t *testing.T) {
		client, err := p.Client()

		// ns := p.Namespace()

		g.Expect(err).To(gomega.BeNil())
		// g.Expect(ns).To(gomega.Equal(testNs))
		g.Expect(client).NotTo(gomega.BeNil())
	})

	t.Run("RunE", func(t *testing.T) {
		err := r.RunE(testCmd, []string{})

		g.Expect(err).To(gomega.BeNil())
	})
}
