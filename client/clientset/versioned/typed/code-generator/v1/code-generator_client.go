package v1

import (
	v1 "github.com/mvladev/code-generation/apis/example/v1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	"github.com/mvladev/code-generation/client/clientset/versioned/scheme"
)


type Code-generatorV1Interface interface {
    RESTClient() rest.Interface
     TestTypesGetter
    
}

// Code-generatorV1Client is used to interact with features provided by the code-generator.k8s.io group.
type Code-generatorV1Client struct {
	restClient rest.Interface
}

func (c *Code-generatorV1Client) TestTypes(namespace string) TestTypeInterface {
	return newTestTypes(c, namespace)
}

// NewForConfig creates a new Code-generatorV1Client for the given config.
func NewForConfig(c *rest.Config) (*Code-generatorV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Code-generatorV1Client{client}, nil
}

// NewForConfigOrDie creates a new Code-generatorV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Code-generatorV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new Code-generatorV1Client for the given RESTClient.
func New(c rest.Interface) *Code-generatorV1Client {
	return &Code-generatorV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion =  &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *Code-generatorV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
