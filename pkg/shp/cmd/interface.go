package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"
)

type SubCommand interface {
	Cmd() *cobra.Command
	Complete(client dynamic.Interface, ns string, args []string) error
	Validate() error
	Run(client dynamic.Interface, ns string) error
}
