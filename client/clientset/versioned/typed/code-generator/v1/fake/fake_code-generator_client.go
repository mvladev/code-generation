package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1"
)


type FakeCode-generatorV1 struct {
	*testing.Fake
}

func (c *FakeCode-generatorV1) TestTypes(namespace string) v1.TestTypeInterface {
	return &FakeTestTypes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCode-generatorV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
