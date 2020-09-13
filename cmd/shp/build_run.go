package main

import (
	"fmt"

	"github.com/otaviof/shp/pkg/shp"
	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/spf13/cobra"
)

// buildRunCmd "build-run" sub-command declaration.
var buildRunCmd = &cobra.Command{
	Use:          "build-run <verb> [options]",
	Short:        "",
	SilenceUsage: true,
	RunE:         runBuildRunCmd,
}

// buildRunSpec spec pointer used to store command-line flag results.
var buildRunSpec *buildv1alpha1.BuildRunSpec

// init prepare command-line flags and linking sub-command with primary instance.
func init() {
	flags := buildRunCmd.PersistentFlags()
	buildRunSpec = shp.BuildRunSpecFlags(flags)
	rootCmd.AddCommand(buildRunCmd)
}

// printBuildRunSpec simple method to display which values are in use for build-run spec, employed
// to store the command-line flags.
func printBuildRunSpec() {
	fmt.Printf("spec.BuildRef:            '%#v'\n", buildRunSpec.BuildRef)
	fmt.Printf("spec.ServiceAccount.Name: '%#v'\n", *buildRunSpec.ServiceAccount.Name)
	fmt.Printf("spec.ServiceAccount:      '%#v'\n", buildRunSpec.ServiceAccount)
	fmt.Printf("spec.Timeout:             '%#v'\n", buildRunSpec.Timeout)
	fmt.Printf("spec.Output.SecretRef:    '%#v'\n", buildRunSpec.Output.SecretRef)
	fmt.Printf("spec.Output:              '%#v'\n", buildRunSpec.Output)
}

// runBuildRunCmd primary logic of the build-run sub-command.
func runBuildRunCmd(cmd *cobra.Command, args []string) error {
	printBuildRunSpec()

	if len(args) != 2 {
		return fmt.Errorf("two arguments are expected, verb and name, informed '%v'", args)
	}

	s, err := newSHP()
	if err != nil {
		return err
	}

	verb := args[0]
	name := args[1]
	switch verb {
	case "create":
		if err = s.BuildRun().Create(name, buildRunSpec); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown verb '%s'", verb)
	}
	return nil
}
