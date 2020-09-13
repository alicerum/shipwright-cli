package main

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestMain_defaultKubeconfigPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	kubeconfigPath := defaultKubeconfigPath()
	expect := g.Expect(kubeconfigPath)

	expect.NotTo(gomega.BeEmpty(), "should have a default kubeconfig path")
	expect.To(gomega.MatchRegexp(`^\/.*?\/%s`, kubeconfigRelativePath), "should match path")
}
