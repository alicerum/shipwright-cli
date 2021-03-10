package resource

import (
	"context"

	"github.com/shipwright-io/cli/pkg/shp/params"
	"github.com/shipwright-io/cli/pkg/shp/util"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Resource struct {
	gv       schema.GroupVersion
	kind     string
	resource string

	params *params.Params

	resourceInterface dynamic.ResourceInterface
}

func (r *Resource) getResourceInterface() (dynamic.ResourceInterface, error) {
	if r.resourceInterface != nil {
		return r.resourceInterface, nil
	}

	client, err := r.params.Client()
	if err != nil {
		return nil, err
	}

	r.resourceInterface = client.Resource(r.gv.WithResource(r.resource)).Namespace(r.params.Namespace())
	return r.resourceInterface, nil
}

func (r *Resource) Create(ctx context.Context, name string, obj interface{}) error {
	ri, err := r.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.CreateObject(ctx, ri, name, r.gv.WithKind(r.kind), obj)
}

func (r *Resource) Delete(ctx context.Context, name string) error {
	ri, err := r.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.DeleteObject(ctx, ri, name, r.gv.WithResource(r.resource))
}

func (r *Resource) List(ctx context.Context, result interface{}) error {
	return r.ListWithOptions(ctx, result, v1.ListOptions{})
}

func (r *Resource) ListWithOptions(ctx context.Context, result interface{}, options v1.ListOptions) error {
	ri, err := r.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.ListObjectWithOptions(ctx, ri, result, options)
}

func (r *Resource) Get(ctx context.Context, name string, result interface{}) error {
	ri, err := r.getResourceInterface()
	if err != nil {
		return nil
	}

	return util.GetObject(ctx, ri, name, result)
}

func NewResource(p *params.Params, gv schema.GroupVersion, kind, resource string) *Resource {
	r := &Resource{
		gv:       gv,
		kind:     kind,
		resource: resource,
		params:   p,
	}

	return r
}
