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

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlertProviders implements AlertProviderInterface
type FakeAlertProviders struct {
	Fake *FakeFlaggerV1beta1
	ns   string
}

var alertprovidersResource = v1beta1.SchemeGroupVersion.WithResource("alertproviders")

var alertprovidersKind = v1beta1.SchemeGroupVersion.WithKind("AlertProvider")

// Get takes name of the alertProvider, and returns the corresponding alertProvider object, and an error if there is any.
func (c *FakeAlertProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AlertProvider, err error) {
	emptyResult := &v1beta1.AlertProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(alertprovidersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.AlertProvider), err
}

// List takes label and field selectors, and returns the list of AlertProviders that match those selectors.
func (c *FakeAlertProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AlertProviderList, err error) {
	emptyResult := &v1beta1.AlertProviderList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(alertprovidersResource, alertprovidersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AlertProviderList{ListMeta: obj.(*v1beta1.AlertProviderList).ListMeta}
	for _, item := range obj.(*v1beta1.AlertProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertProviders.
func (c *FakeAlertProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(alertprovidersResource, c.ns, opts))

}

// Create takes the representation of a alertProvider and creates it.  Returns the server's representation of the alertProvider, and an error, if there is any.
func (c *FakeAlertProviders) Create(ctx context.Context, alertProvider *v1beta1.AlertProvider, opts v1.CreateOptions) (result *v1beta1.AlertProvider, err error) {
	emptyResult := &v1beta1.AlertProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(alertprovidersResource, c.ns, alertProvider, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.AlertProvider), err
}

// Update takes the representation of a alertProvider and updates it. Returns the server's representation of the alertProvider, and an error, if there is any.
func (c *FakeAlertProviders) Update(ctx context.Context, alertProvider *v1beta1.AlertProvider, opts v1.UpdateOptions) (result *v1beta1.AlertProvider, err error) {
	emptyResult := &v1beta1.AlertProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(alertprovidersResource, c.ns, alertProvider, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.AlertProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertProviders) UpdateStatus(ctx context.Context, alertProvider *v1beta1.AlertProvider, opts v1.UpdateOptions) (result *v1beta1.AlertProvider, err error) {
	emptyResult := &v1beta1.AlertProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(alertprovidersResource, "status", c.ns, alertProvider, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.AlertProvider), err
}

// Delete takes name of the alertProvider and deletes it. Returns an error if one occurs.
func (c *FakeAlertProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alertprovidersResource, c.ns, name, opts), &v1beta1.AlertProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(alertprovidersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AlertProviderList{})
	return err
}

// Patch applies the patch and returns the patched alertProvider.
func (c *FakeAlertProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AlertProvider, err error) {
	emptyResult := &v1beta1.AlertProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(alertprovidersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.AlertProvider), err
}
