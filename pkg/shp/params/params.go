package params

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ShipwrightParams struct {
	client         dynamic.Interface
	kubeConfigPath string
	kubeContext    string
	namespace      string
}

func (p *ShipwrightParams) config() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if p.kubeConfigPath != "" {
		loadingRules.ExplicitPath = p.kubeConfigPath
	}

	configOverrides := &clientcmd.ConfigOverrides{}
	if p.kubeContext != "" {
		configOverrides.CurrentContext = p.kubeContext
	}

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	if p.namespace == "" {
		namespace, _, err := kubeConfig.Namespace()
		if err != nil {
			return nil, errors.Wrap(err, "Could not get namespace from KubeConfig")
		}
		p.namespace = namespace
	}

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse KubeConfig")
	}

	return config, nil
}

func (p *ShipwrightParams) SetKubeConfigPath(kubeConfigPath string) {
	p.kubeConfigPath = kubeConfigPath
}

func (p *ShipwrightParams) SetKubeContext(kubeContext string) {
	p.kubeContext = kubeContext
}

func (p *ShipwrightParams) Client() (dynamic.Interface, error) {
	if p.client != nil {
		return p.client, nil
	}

	config, err := p.config()
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
	return p.namespace
}

func NewParams(configPath string, namespace string, context string) (Params, error) {
	p := &ShipwrightParams{
		namespace:      namespace,
		kubeContext:    context,
		kubeConfigPath: configPath,
	}

	// initialize clients to be used immediately
	// along with struct fielst like Namespace
	_, err := p.Client()
	return p, err
}
