package util

import (
	"testing"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"

	"github.com/onsi/gomega"
	"github.com/shipwright-io/cli/test/stub"
)

func TestUtil_ToUnstructured(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	name := "test"
	kind := "BuildRun"
	gvk := buildv1alpha1.SchemeBuilder.GroupVersion.WithKind(kind)
	br := stub.BuildRunEmpty()

	u, err := toUnstructured(name, gvk, &br)

	g.Expect(err).To(gomega.BeNil())
	g.Expect(u.GetName()).To(gomega.Equal(name))
	g.Expect(u.GetKind()).To(gomega.Equal(kind))
}
