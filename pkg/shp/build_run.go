package shp

import (
	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
)

// BuildRun represents Shipwight BuildRun resource, and the verbs that can be executed against
// this resource.
type BuildRun struct {
	client dynamic.ResourceInterface
}

// BuildRunKind BuildRun API resource kind.
const BuildRunKind = "BuildRun"

// Create a BuildRun resource with informed name and spec.
func (b *BuildRun) Create(name string, spec *buildv1alpha1.BuildRunSpec) error {
	gvk := buildv1alpha1.SchemeGroupVersion.WithKind(BuildRunKind)
	u, err := ToUnstructured(name, gvk, &buildv1alpha1.BuildRun{Spec: *spec})
	if err != nil {
		return err
	}

	_, err = b.client.Create(u, metav1.CreateOptions{})
	return err
}

// NewBuildRun instantiate the BuildRun handler.
func NewBuildRun(client dynamic.ResourceInterface) *BuildRun {
	return &BuildRun{client: client}
}
