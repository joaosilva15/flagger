/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTraefikServices implements TraefikServiceInterface
type FakeTraefikServices struct {
	Fake *FakeTraefikioV1alpha1
	ns   string
}

var traefikservicesResource = v1alpha1.SchemeGroupVersion.WithResource("traefikservices")

var traefikservicesKind = v1alpha1.SchemeGroupVersion.WithKind("TraefikService")

// Get takes name of the traefikService, and returns the corresponding traefikService object, and an error if there is any.
func (c *FakeTraefikServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TraefikService, err error) {
	emptyResult := &v1alpha1.TraefikService{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(traefikservicesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TraefikService), err
}

// List takes label and field selectors, and returns the list of TraefikServices that match those selectors.
func (c *FakeTraefikServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TraefikServiceList, err error) {
	emptyResult := &v1alpha1.TraefikServiceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(traefikservicesResource, traefikservicesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TraefikServiceList{ListMeta: obj.(*v1alpha1.TraefikServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.TraefikServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested traefikServices.
func (c *FakeTraefikServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(traefikservicesResource, c.ns, opts))

}

// Create takes the representation of a traefikService and creates it.  Returns the server's representation of the traefikService, and an error, if there is any.
func (c *FakeTraefikServices) Create(ctx context.Context, traefikService *v1alpha1.TraefikService, opts v1.CreateOptions) (result *v1alpha1.TraefikService, err error) {
	emptyResult := &v1alpha1.TraefikService{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(traefikservicesResource, c.ns, traefikService, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TraefikService), err
}

// Update takes the representation of a traefikService and updates it. Returns the server's representation of the traefikService, and an error, if there is any.
func (c *FakeTraefikServices) Update(ctx context.Context, traefikService *v1alpha1.TraefikService, opts v1.UpdateOptions) (result *v1alpha1.TraefikService, err error) {
	emptyResult := &v1alpha1.TraefikService{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(traefikservicesResource, c.ns, traefikService, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TraefikService), err
}

// Delete takes name of the traefikService and deletes it. Returns an error if one occurs.
func (c *FakeTraefikServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(traefikservicesResource, c.ns, name, opts), &v1alpha1.TraefikService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTraefikServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(traefikservicesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TraefikServiceList{})
	return err
}

// Patch applies the patch and returns the patched traefikService.
func (c *FakeTraefikServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TraefikService, err error) {
	emptyResult := &v1alpha1.TraefikService{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(traefikservicesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.TraefikService), err
}
