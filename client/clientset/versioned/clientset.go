package versioned

import (
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	code-generatorv1 "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1"
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
)


type Interface interface {
	Discovery() discovery.DiscoveryInterface
    Code-generatorV1() code-generatorv1.Code-generatorV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Code-generator() code-generatorv1.Code-generatorV1Interface
	
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
    code-generatorV1 *code-generatorv1.Code-generatorV1Client
    
}

// Code-generatorV1 retrieves the Code-generatorV1Client
func (c *Clientset) Code-generatorV1() code-generatorv1.Code-generatorV1Interface {
	return c.code-generatorV1
}

// Deprecated: Code-generator retrieves the default version of Code-generatorClient.
// Please explicitly pick a version.
func (c *Clientset) Code-generator() code-generatorv1.Code-generatorV1Interface {
	return c.code-generatorV1
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
    cs.code-generatorV1, err =code-generatorv1.NewForConfig(&configShallowCopy)
	if err!=nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err!=nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
    cs.code-generatorV1 =code-generatorv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
    cs.code-generatorV1 =code-generatorv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
