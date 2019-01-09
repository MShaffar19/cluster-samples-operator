// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	glog "github.com/golang/glog"
	samplesresourcev1alpha1 "github.com/openshift/cluster-samples-operator/pkg/generated/clientset/versioned/typed/samplesresource/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	SamplesresourceV1alpha1() samplesresourcev1alpha1.SamplesresourceV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Samplesresource() samplesresourcev1alpha1.SamplesresourceV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	samplesresourceV1alpha1 *samplesresourcev1alpha1.SamplesresourceV1alpha1Client
}

// SamplesresourceV1alpha1 retrieves the SamplesresourceV1alpha1Client
func (c *Clientset) SamplesresourceV1alpha1() samplesresourcev1alpha1.SamplesresourceV1alpha1Interface {
	return c.samplesresourceV1alpha1
}

// Deprecated: Samplesresource retrieves the default version of SamplesresourceClient.
// Please explicitly pick a version.
func (c *Clientset) Samplesresource() samplesresourcev1alpha1.SamplesresourceV1alpha1Interface {
	return c.samplesresourceV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.samplesresourceV1alpha1, err = samplesresourcev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.samplesresourceV1alpha1 = samplesresourcev1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.samplesresourceV1alpha1 = samplesresourcev1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
