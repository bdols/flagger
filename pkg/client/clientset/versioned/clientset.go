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

package versioned

import (
	"fmt"
	"net/http"

	apisixv2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/apisix/v2"
	appmeshv1beta1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta1"
	appmeshv1beta2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta2"
	flaggerv1beta1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/flagger/v1beta1"
	gatewayv1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/gateway/v1"
	gatewayapiv1alpha2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/gatewayapi/v1alpha2"
	gatewayapiv1beta1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/gatewayapi/v1beta1"
	gloov1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/gloo/v1"
	networkingv1alpha3 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	kedav1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/keda/v1alpha1"
	kumav1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/kuma/v1alpha1"
	projectcontourv1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/projectcontour/v1"
	splitv1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha1"
	splitv1alpha2 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha2"
	splitv1alpha3 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha3"
	traefikv1alpha1 "github.com/fluxcd/flagger/pkg/client/clientset/versioned/typed/traefik/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApisixV2() apisixv2.ApisixV2Interface
	AppmeshV1beta2() appmeshv1beta2.AppmeshV1beta2Interface
	AppmeshV1beta1() appmeshv1beta1.AppmeshV1beta1Interface
	FlaggerV1beta1() flaggerv1beta1.FlaggerV1beta1Interface
	GatewayV1() gatewayv1.GatewayV1Interface
	GatewayapiV1alpha2() gatewayapiv1alpha2.GatewayapiV1alpha2Interface
	GatewayapiV1beta1() gatewayapiv1beta1.GatewayapiV1beta1Interface
	GlooV1() gloov1.GlooV1Interface
	NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface
	KedaV1alpha1() kedav1alpha1.KedaV1alpha1Interface
	KumaV1alpha1() kumav1alpha1.KumaV1alpha1Interface
	ProjectcontourV1() projectcontourv1.ProjectcontourV1Interface
	SplitV1alpha1() splitv1alpha1.SplitV1alpha1Interface
	SplitV1alpha2() splitv1alpha2.SplitV1alpha2Interface
	SplitV1alpha3() splitv1alpha3.SplitV1alpha3Interface
	TraefikV1alpha1() traefikv1alpha1.TraefikV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	apisixV2           *apisixv2.ApisixV2Client
	appmeshV1beta2     *appmeshv1beta2.AppmeshV1beta2Client
	appmeshV1beta1     *appmeshv1beta1.AppmeshV1beta1Client
	flaggerV1beta1     *flaggerv1beta1.FlaggerV1beta1Client
	gatewayV1          *gatewayv1.GatewayV1Client
	gatewayapiV1alpha2 *gatewayapiv1alpha2.GatewayapiV1alpha2Client
	gatewayapiV1beta1  *gatewayapiv1beta1.GatewayapiV1beta1Client
	glooV1             *gloov1.GlooV1Client
	networkingV1alpha3 *networkingv1alpha3.NetworkingV1alpha3Client
	kedaV1alpha1       *kedav1alpha1.KedaV1alpha1Client
	kumaV1alpha1       *kumav1alpha1.KumaV1alpha1Client
	projectcontourV1   *projectcontourv1.ProjectcontourV1Client
	splitV1alpha1      *splitv1alpha1.SplitV1alpha1Client
	splitV1alpha2      *splitv1alpha2.SplitV1alpha2Client
	splitV1alpha3      *splitv1alpha3.SplitV1alpha3Client
	traefikV1alpha1    *traefikv1alpha1.TraefikV1alpha1Client
}

// ApisixV2 retrieves the ApisixV2Client
func (c *Clientset) ApisixV2() apisixv2.ApisixV2Interface {
	return c.apisixV2
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

// GatewayV1 retrieves the GatewayV1Client
func (c *Clientset) GatewayV1() gatewayv1.GatewayV1Interface {
	return c.gatewayV1
}

// GatewayapiV1alpha2 retrieves the GatewayapiV1alpha2Client
func (c *Clientset) GatewayapiV1alpha2() gatewayapiv1alpha2.GatewayapiV1alpha2Interface {
	return c.gatewayapiV1alpha2
}

// GatewayapiV1beta1 retrieves the GatewayapiV1beta1Client
func (c *Clientset) GatewayapiV1beta1() gatewayapiv1beta1.GatewayapiV1beta1Interface {
	return c.gatewayapiV1beta1
}

// GlooV1 retrieves the GlooV1Client
func (c *Clientset) GlooV1() gloov1.GlooV1Interface {
	return c.glooV1
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// KedaV1alpha1 retrieves the KedaV1alpha1Client
func (c *Clientset) KedaV1alpha1() kedav1alpha1.KedaV1alpha1Interface {
	return c.kedaV1alpha1
}

// KumaV1alpha1 retrieves the KumaV1alpha1Client
func (c *Clientset) KumaV1alpha1() kumav1alpha1.KumaV1alpha1Interface {
	return c.kumaV1alpha1
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

// SplitV1alpha3 retrieves the SplitV1alpha3Client
func (c *Clientset) SplitV1alpha3() splitv1alpha3.SplitV1alpha3Interface {
	return c.splitV1alpha3
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
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.apisixV2, err = apisixv2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appmeshV1beta2, err = appmeshv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appmeshV1beta1, err = appmeshv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flaggerV1beta1, err = flaggerv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.gatewayV1, err = gatewayv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.gatewayapiV1alpha2, err = gatewayapiv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.gatewayapiV1beta1, err = gatewayapiv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.glooV1, err = gloov1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1alpha3, err = networkingv1alpha3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kedaV1alpha1, err = kedav1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kumaV1alpha1, err = kumav1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.projectcontourV1, err = projectcontourv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.splitV1alpha1, err = splitv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.splitV1alpha2, err = splitv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.splitV1alpha3, err = splitv1alpha3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.traefikV1alpha1, err = traefikv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.apisixV2 = apisixv2.New(c)
	cs.appmeshV1beta2 = appmeshv1beta2.New(c)
	cs.appmeshV1beta1 = appmeshv1beta1.New(c)
	cs.flaggerV1beta1 = flaggerv1beta1.New(c)
	cs.gatewayV1 = gatewayv1.New(c)
	cs.gatewayapiV1alpha2 = gatewayapiv1alpha2.New(c)
	cs.gatewayapiV1beta1 = gatewayapiv1beta1.New(c)
	cs.glooV1 = gloov1.New(c)
	cs.networkingV1alpha3 = networkingv1alpha3.New(c)
	cs.kedaV1alpha1 = kedav1alpha1.New(c)
	cs.kumaV1alpha1 = kumav1alpha1.New(c)
	cs.projectcontourV1 = projectcontourv1.New(c)
	cs.splitV1alpha1 = splitv1alpha1.New(c)
	cs.splitV1alpha2 = splitv1alpha2.New(c)
	cs.splitV1alpha3 = splitv1alpha3.New(c)
	cs.traefikV1alpha1 = traefikv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
