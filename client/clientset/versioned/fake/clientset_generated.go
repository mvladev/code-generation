package fake

import (
	fakecode-generatorv1 "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/fake"
	clientset "github.com/mvladev/code-generation/client/clientset/versioned"
	"k8s.io/client-go/testing"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	code-generatorv1 "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1"
)


// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// Code-generatorV1 retrieves the Code-generatorV1Client
func (c *Clientset) Code-generatorV1() code-generatorv1.Code-generatorV1Interface {
	return &fakecode-generatorv1.FakeCode-generatorV1{Fake: &c.Fake}
}

// Code-generator retrieves the Code-generatorV1Client
func (c *Clientset) Code-generator() code-generatorv1.Code-generatorV1Interface {
	return &fakecode-generatorv1.FakeCode-generatorV1{Fake: &c.Fake}
}
