package params

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
)

type ShipwrightParams struct {
	client dynamic.Interface

	configFlags *genericclioptions.ConfigFlags
}

func (p *ShipwrightParams) AddFlags(flags *pflag.FlagSet) {
	p.configFlags.AddFlags(flags)
}

func (p *ShipwrightParams) Client() (dynamic.Interface, error) {
	if p.client != nil {
		return p.client, nil
	}

	config, err := p.configFlags.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	dynamic, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create Dynamic client")
	}

	p.client = dynamic

	return p.client, nil
}

func (p *ShipwrightParams) Namespace() string {
	return *p.configFlags.Namespace
}

func NewParams() Params {
	p := &ShipwrightParams{}
	p.configFlags = genericclioptions.NewConfigFlags(true)

	return p
}
