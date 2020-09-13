package shp

import (
	"context"

	"k8s.io/client-go/dynamic"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
)

// SHP represents the primary type for this command-line project, it links together all components
// needed, and share them with inner types.
type SHP struct {
	ctx       context.Context   // shared context
	namespace string            // namespace name
	client    dynamic.Interface // kubernetes api client
}

// BuildRun handles requests against BuildRun CRD.
func (s *SHP) BuildRun() *BuildRun {
	gvr := buildv1alpha1.SchemeGroupVersion.WithResource("buildruns")
	resourceClient := s.client.Resource(gvr).Namespace(s.namespace)
	return NewBuildRun(resourceClient)
}

// NewSHP instantiate SHP.
func NewSHP(ctx context.Context, namespace string, client dynamic.Interface) *SHP {
	return &SHP{ctx: ctx, namespace: namespace, client: client}
}
