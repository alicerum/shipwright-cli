package build

import (
	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"

	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/resource"
)

// GetBuildResource returns dynamic client resource
// for working with Build objects in kubernetes
func GetBuildResource(p *params.Params) *resource.Resource {
	return resource.NewResource(
		p,
		buildv1alpha1.SchemeBuilder.GroupVersion,
		"Build",
		"builds",
	)
}
