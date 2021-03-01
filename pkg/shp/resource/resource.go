package resource

import (
	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/util"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type ShpResource struct {
	gv       schema.GroupVersion
	kind     string
	resource string

	params params.Params
}

func (sr *ShpResource) getResourceInterface() (dynamic.ResourceInterface, error) {
	client, err := sr.params.Client()
	if err != nil {
		return nil, err
	}

	ri := client.Resource(sr.gv.WithResource(sr.resource)).Namespace(sr.params.Namespace())
	return ri, nil
}

func (sr *ShpResource) Create(name string, obj interface{}) error {
	ri, err := sr.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.CreateObject(ri, name, sr.gv.WithKind(sr.kind), obj)
}

func (sr *ShpResource) Delete(name string) error {
	ri, err := sr.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.DeleteObject(ri, name, sr.gv.WithResource(sr.resource))
}

func (sr *ShpResource) List(result interface{}) error {
	ri, err := sr.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.ListObject(ri, result)
}

func (sr *ShpResource) Get(name string, result interface{}) error {
	ri, err := sr.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.GetObject(ri, name, result)
}

func NewShpResource(p params.Params, gv schema.GroupVersion, kind, resource string) *ShpResource {
	sr := &ShpResource{
		gv:       gv,
		kind:     kind,
		resource: resource,
		params:   p,
	}

	return sr
}
