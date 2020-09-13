package main

import (
	"context"
	"path"

	"github.com/otaviof/shp/pkg/shp"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const kubeconfigRelativePath = ".kube/config"

// defaultKubeconfigPath returns the path to kubeconfig under home directory, when home directory
// can't be asserted, it simply returns empty.
func defaultKubeconfigPath() string {
	home := homedir.HomeDir()
	if home == "" {
		return ""
	}
	return path.Join(home, kubeconfigRelativePath)
}

// newDynamicClient instantiate a Kubernetes dynamic client with informed kubeconfig path and
// Kubernetes context.
func newDynamicClient(kubeconfigPath string, kubeContext string) (dynamic.Interface, error) {
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{
			ExplicitPath: kubeconfigPath,
		},
		&clientcmd.ConfigOverrides{
			CurrentContext: kubeContext,
		},
	)

	// when namespace is not informed, it will uses the default namespace configured in kubeconfig
	if namespace == "" {
		var err error
		namespace, _, err = clientConfig.Namespace()
		if err != nil {
			return nil, err
		}
	}

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(restConfig)
}

// newSHP instantiate SHP.
func newSHP() (*shp.SHP, error) {
	client, err := newDynamicClient(kubeconfig, kubeContext)
	if err != nil {
		return nil, err
	}
	return shp.NewSHP(context.TODO(), namespace, client), nil
}
