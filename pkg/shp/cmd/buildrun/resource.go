package buildrun

import (
	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/resource"
)

func GetBuildRunResource(p params.Params) *resource.ShpResource {
	return resource.NewShpResource(
		p,
		buildv1alpha1.SchemeGroupVersion,
		"BuildRun",
		"buildruns",
	)
}
