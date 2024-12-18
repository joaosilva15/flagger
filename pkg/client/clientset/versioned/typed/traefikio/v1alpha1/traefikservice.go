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

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/traefikio/v1alpha1"
	scheme "github.com/fluxcd/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TraefikServicesGetter has a method to return a TraefikServiceInterface.
// A group's client should implement this interface.
type TraefikServicesGetter interface {
	TraefikServices(namespace string) TraefikServiceInterface
}

// TraefikServiceInterface has methods to work with TraefikService resources.
type TraefikServiceInterface interface {
	Create(ctx context.Context, traefikService *v1alpha1.TraefikService, opts v1.CreateOptions) (*v1alpha1.TraefikService, error)
	Update(ctx context.Context, traefikService *v1alpha1.TraefikService, opts v1.UpdateOptions) (*v1alpha1.TraefikService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TraefikService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TraefikServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TraefikService, err error)
	TraefikServiceExpansion
}

// traefikServices implements TraefikServiceInterface
type traefikServices struct {
	*gentype.ClientWithList[*v1alpha1.TraefikService, *v1alpha1.TraefikServiceList]
}

// newTraefikServices returns a TraefikServices
func newTraefikServices(c *TraefikioV1alpha1Client, namespace string) *traefikServices {
	return &traefikServices{
		gentype.NewClientWithList[*v1alpha1.TraefikService, *v1alpha1.TraefikServiceList](
			"traefikservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.TraefikService { return &v1alpha1.TraefikService{} },
			func() *v1alpha1.TraefikServiceList { return &v1alpha1.TraefikServiceList{} }),
	}
}
