package shp

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ToUnstructured converts informed object to unstructured.
func ToUnstructured(
	name string,
	gvk schema.GroupVersionKind,
	obj interface{},
) (*unstructured.Unstructured, error) {
	data, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: data}
	u.SetName(name)
	u.SetGroupVersionKind(gvk)
	return u, nil
}
