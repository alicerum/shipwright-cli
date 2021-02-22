package params

import (
	"k8s.io/client-go/dynamic"
)

type Params interface {
	Client() (dynamic.Interface, error)

	Namespace() string

	SetKubeConfigPath(kubeConfigPath string)
	SetKubeContext(kubeContext string)
}
