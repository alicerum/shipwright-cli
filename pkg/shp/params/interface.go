package params

import (
	"github.com/spf13/pflag"
	"k8s.io/client-go/dynamic"
)

// Params is an interface that provides
// functions for program to intaract with its parameters
type Params interface {
	Client() (dynamic.Interface, error)

	Namespace() string

	AddFlags(f *pflag.FlagSet)
}
