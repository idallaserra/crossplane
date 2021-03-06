/*
Copyright 2019 The Crossplane Authors.

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

	apiextensionsv1 "github.com/crossplane/crossplane/internal/client/clientset/versioned/typed/apiextensions/v1"
	pkgv1 "github.com/crossplane/crossplane/internal/client/clientset/versioned/typed/pkg/v1"
	pkgv1alpha1 "github.com/crossplane/crossplane/internal/client/clientset/versioned/typed/pkg/v1alpha1"
	pkgv1beta1 "github.com/crossplane/crossplane/internal/client/clientset/versioned/typed/pkg/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApiextensionsV1() apiextensionsv1.ApiextensionsV1Interface
	PkgV1alpha1() pkgv1alpha1.PkgV1alpha1Interface
	PkgV1beta1() pkgv1beta1.PkgV1beta1Interface
	PkgV1() pkgv1.PkgV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	apiextensionsV1 *apiextensionsv1.ApiextensionsV1Client
	pkgV1alpha1     *pkgv1alpha1.PkgV1alpha1Client
	pkgV1beta1      *pkgv1beta1.PkgV1beta1Client
	pkgV1           *pkgv1.PkgV1Client
}

// ApiextensionsV1 retrieves the ApiextensionsV1Client
func (c *Clientset) ApiextensionsV1() apiextensionsv1.ApiextensionsV1Interface {
	return c.apiextensionsV1
}

// PkgV1alpha1 retrieves the PkgV1alpha1Client
func (c *Clientset) PkgV1alpha1() pkgv1alpha1.PkgV1alpha1Interface {
	return c.pkgV1alpha1
}

// PkgV1beta1 retrieves the PkgV1beta1Client
func (c *Clientset) PkgV1beta1() pkgv1beta1.PkgV1beta1Interface {
	return c.pkgV1beta1
}

// PkgV1 retrieves the PkgV1Client
func (c *Clientset) PkgV1() pkgv1.PkgV1Interface {
	return c.pkgV1
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
	cs.apiextensionsV1, err = apiextensionsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.pkgV1alpha1, err = pkgv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.pkgV1beta1, err = pkgv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.pkgV1, err = pkgv1.NewForConfig(&configShallowCopy)
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
	cs.apiextensionsV1 = apiextensionsv1.NewForConfigOrDie(c)
	cs.pkgV1alpha1 = pkgv1alpha1.NewForConfigOrDie(c)
	cs.pkgV1beta1 = pkgv1beta1.NewForConfigOrDie(c)
	cs.pkgV1 = pkgv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.apiextensionsV1 = apiextensionsv1.New(c)
	cs.pkgV1alpha1 = pkgv1alpha1.New(c)
	cs.pkgV1beta1 = pkgv1beta1.New(c)
	cs.pkgV1 = pkgv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
