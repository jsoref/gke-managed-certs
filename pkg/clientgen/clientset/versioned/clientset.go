/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

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

	networkingv1 "github.com/GoogleCloudPlatform/gke-managed-certs/pkg/clientgen/clientset/versioned/typed/networking.gke.io/v1"
	networkingv1beta1 "github.com/GoogleCloudPlatform/gke-managed-certs/pkg/clientgen/clientset/versioned/typed/networking.gke.io/v1beta1"
	networkingv1beta2 "github.com/GoogleCloudPlatform/gke-managed-certs/pkg/clientgen/clientset/versioned/typed/networking.gke.io/v1beta2"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface
	NetworkingV1beta2() networkingv1beta2.NetworkingV1beta2Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	networkingV1beta1 *networkingv1beta1.NetworkingV1beta1Client
	networkingV1beta2 *networkingv1beta2.NetworkingV1beta2Client
	networkingV1      *networkingv1.NetworkingV1Client
}

// NetworkingV1beta1 retrieves the NetworkingV1beta1Client
func (c *Clientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return c.networkingV1beta1
}

// NetworkingV1beta2 retrieves the NetworkingV1beta2Client
func (c *Clientset) NetworkingV1beta2() networkingv1beta2.NetworkingV1beta2Interface {
	return c.networkingV1beta2
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
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
	cs.networkingV1beta1, err = networkingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1beta2, err = networkingv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfig(&configShallowCopy)
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
	cs.networkingV1beta1 = networkingv1beta1.NewForConfigOrDie(c)
	cs.networkingV1beta2 = networkingv1beta2.NewForConfigOrDie(c)
	cs.networkingV1 = networkingv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.networkingV1beta1 = networkingv1beta1.New(c)
	cs.networkingV1beta2 = networkingv1beta2.New(c)
	cs.networkingV1 = networkingv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
