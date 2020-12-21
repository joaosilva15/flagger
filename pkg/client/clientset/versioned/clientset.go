/*
Copyright The Flagger Authors.

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

package versioned

import (
	"fmt"

	appmeshv1beta1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta1"
	appmeshv1beta2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta2"
	flaggerv1beta1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/flagger/v1beta1"
	gloov1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/gloo/v1"
	networkingv1alpha3 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	projectcontourv1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/projectcontour/v1"
	splitv1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha1"
	splitv1alpha2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha2"
	traefikv1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/traefik/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AppmeshV1beta2() appmeshv1beta2.AppmeshV1beta2Interface
	AppmeshV1beta1() appmeshv1beta1.AppmeshV1beta1Interface
	FlaggerV1beta1() flaggerv1beta1.FlaggerV1beta1Interface
	GlooV1() gloov1.GlooV1Interface
	NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface
	ProjectcontourV1() projectcontourv1.ProjectcontourV1Interface
	SplitV1alpha1() splitv1alpha1.SplitV1alpha1Interface
	SplitV1alpha2() splitv1alpha2.SplitV1alpha2Interface
	TraefikV1alpha1() traefikv1alpha1.TraefikV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	appmeshV1beta2     *appmeshv1beta2.AppmeshV1beta2Client
	appmeshV1beta1     *appmeshv1beta1.AppmeshV1beta1Client
	flaggerV1beta1     *flaggerv1beta1.FlaggerV1beta1Client
	glooV1             *gloov1.GlooV1Client
	networkingV1alpha3 *networkingv1alpha3.NetworkingV1alpha3Client
	projectcontourV1   *projectcontourv1.ProjectcontourV1Client
	splitV1alpha1      *splitv1alpha1.SplitV1alpha1Client
	splitV1alpha2      *splitv1alpha2.SplitV1alpha2Client
	traefikV1alpha1    *traefikv1alpha1.TraefikV1alpha1Client
}

// AppmeshV1beta2 retrieves the AppmeshV1beta2Client
func (c *Clientset) AppmeshV1beta2() appmeshv1beta2.AppmeshV1beta2Interface {
	return c.appmeshV1beta2
}

// AppmeshV1beta1 retrieves the AppmeshV1beta1Client
func (c *Clientset) AppmeshV1beta1() appmeshv1beta1.AppmeshV1beta1Interface {
	return c.appmeshV1beta1
}

// FlaggerV1beta1 retrieves the FlaggerV1beta1Client
func (c *Clientset) FlaggerV1beta1() flaggerv1beta1.FlaggerV1beta1Interface {
	return c.flaggerV1beta1
}

// GlooV1 retrieves the GlooV1Client
func (c *Clientset) GlooV1() gloov1.GlooV1Interface {
	return c.glooV1
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// ProjectcontourV1 retrieves the ProjectcontourV1Client
func (c *Clientset) ProjectcontourV1() projectcontourv1.ProjectcontourV1Interface {
	return c.projectcontourV1
}

// SplitV1alpha1 retrieves the SplitV1alpha1Client
func (c *Clientset) SplitV1alpha1() splitv1alpha1.SplitV1alpha1Interface {
	return c.splitV1alpha1
}

// SplitV1alpha2 retrieves the SplitV1alpha2Client
func (c *Clientset) SplitV1alpha2() splitv1alpha2.SplitV1alpha2Interface {
	return c.splitV1alpha2
}

// TraefikV1alpha1 retrieves the TraefikV1alpha1Client
func (c *Clientset) TraefikV1alpha1() traefikv1alpha1.TraefikV1alpha1Interface {
	return c.traefikV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.appmeshV1beta2, err = appmeshv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appmeshV1beta1, err = appmeshv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.flaggerV1beta1, err = flaggerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.glooV1, err = gloov1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1alpha3, err = networkingv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.projectcontourV1, err = projectcontourv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.splitV1alpha1, err = splitv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.splitV1alpha2, err = splitv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.traefikV1alpha1, err = traefikv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.appmeshV1beta2 = appmeshv1beta2.NewForConfigOrDie(c)
	cs.appmeshV1beta1 = appmeshv1beta1.NewForConfigOrDie(c)
	cs.flaggerV1beta1 = flaggerv1beta1.NewForConfigOrDie(c)
	cs.glooV1 = gloov1.NewForConfigOrDie(c)
	cs.networkingV1alpha3 = networkingv1alpha3.NewForConfigOrDie(c)
	cs.projectcontourV1 = projectcontourv1.NewForConfigOrDie(c)
	cs.splitV1alpha1 = splitv1alpha1.NewForConfigOrDie(c)
	cs.splitV1alpha2 = splitv1alpha2.NewForConfigOrDie(c)
	cs.traefikV1alpha1 = traefikv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.appmeshV1beta2 = appmeshv1beta2.New(c)
	cs.appmeshV1beta1 = appmeshv1beta1.New(c)
	cs.flaggerV1beta1 = flaggerv1beta1.New(c)
	cs.glooV1 = gloov1.New(c)
	cs.networkingV1alpha3 = networkingv1alpha3.New(c)
	cs.projectcontourV1 = projectcontourv1.New(c)
	cs.splitV1alpha1 = splitv1alpha1.New(c)
	cs.splitV1alpha2 = splitv1alpha2.New(c)
	cs.traefikV1alpha1 = traefikv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
