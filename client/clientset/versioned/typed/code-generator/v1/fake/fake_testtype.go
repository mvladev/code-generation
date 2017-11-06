package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	example_v1 "github.com/mvladev/code-generation/apis/example/v1"
)


// FakeTestTypes implements TestTypeInterface
type FakeTestTypes struct {
	Fake *FakeCode-generatorV1
	ns     string
}

var testtypesResource = schema.GroupVersionResource{Group: "code-generator.k8s.io", Version: "v1", Resource: "testtypes"}

var testtypesKind = schema.GroupVersionKind{Group: "code-generator.k8s.io", Version: "v1", Kind: "TestType"}

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *FakeTestTypes) Get(name string, options v1.GetOptions) (result *example_v1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testtypesResource, c.ns, name), &example_v1.TestType{})
		
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.TestType), err
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *FakeTestTypes) List(opts v1.ListOptions) (result *example_v1.TestTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testtypesResource, testtypesKind, c.ns, opts), &example_v1.TestTypeList{})
		
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &example_v1.TestTypeList{}
	for _, item := range obj.(*example_v1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testTypes.
func (c *FakeTestTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testtypesResource, c.ns, opts))
		
}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Create(testType *example_v1.TestType) (result *example_v1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testtypesResource, c.ns, testType), &example_v1.TestType{})
		
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.TestType), err
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Update(testType *example_v1.TestType) (result *example_v1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testtypesResource, c.ns, testType), &example_v1.TestType{})
		
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.TestType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTestTypes) UpdateStatus(testType *example_v1.TestType) (*example_v1.TestType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(testtypesResource, "status", c.ns, testType), &example_v1.TestType{})
		
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.TestType), err
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *FakeTestTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testtypesResource, c.ns, name), &example_v1.TestType{})
		
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testtypesResource, c.ns, listOptions)
	
	_, err := c.Fake.Invokes(action, &example_v1.TestTypeList{})
	return err
}

// Patch applies the patch and returns the patched testType.
func (c *FakeTestTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *example_v1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testtypesResource, c.ns, name, data, subresources... ), &example_v1.TestType{})
		
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.TestType), err
}
