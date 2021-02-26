package params

import (
	"github.com/spf13/pflag"
	"k8s.io/client-go/dynamic"
)

type Params interface {
	Client() (dynamic.Interface, error)

	Namespace() string

	AddFlags(f *pflag.FlagSet)
}
