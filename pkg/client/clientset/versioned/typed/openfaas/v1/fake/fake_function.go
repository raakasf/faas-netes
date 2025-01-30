/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/openfaas/faas-netes/pkg/apis/openfaas/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFunctions implements FunctionInterface
type FakeFunctions struct {
	Fake *FakeOpenfaasV1
	ns   string
}

var functionsResource = v1.SchemeGroupVersion.WithResource("functions")

var functionsKind = v1.SchemeGroupVersion.WithKind("Function")

// Get takes name of the function, and returns the corresponding function object, and an error if there is any.
func (c *FakeFunctions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(functionsResource, c.ns, name), &v1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Function), err
}

// List takes label and field selectors, and returns the list of Functions that match those selectors.
func (c *FakeFunctions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FunctionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(functionsResource, functionsKind, c.ns, opts), &v1.FunctionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.FunctionList{ListMeta: obj.(*v1.FunctionList).ListMeta}
	for _, item := range obj.(*v1.FunctionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested functions.
func (c *FakeFunctions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(functionsResource, c.ns, opts))

}

// Create takes the representation of a function and creates it.  Returns the server's representation of the function, and an error, if there is any.
func (c *FakeFunctions) Create(ctx context.Context, function *v1.Function, opts metav1.CreateOptions) (result *v1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(functionsResource, c.ns, function), &v1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Function), err
}

// Update takes the representation of a function and updates it. Returns the server's representation of the function, and an error, if there is any.
func (c *FakeFunctions) Update(ctx context.Context, function *v1.Function, opts metav1.UpdateOptions) (result *v1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(functionsResource, c.ns, function), &v1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Function), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFunctions) UpdateStatus(ctx context.Context, function *v1.Function, opts metav1.UpdateOptions) (*v1.Function, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(functionsResource, "status", c.ns, function), &v1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Function), err
}

// Delete takes name of the function and deletes it. Returns an error if one occurs.
func (c *FakeFunctions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(functionsResource, c.ns, name, opts), &v1.Function{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFunctions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(functionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.FunctionList{})
	return err
}

// Patch applies the patch and returns the patched function.
func (c *FakeFunctions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(functionsResource, c.ns, name, pt, data, subresources...), &v1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Function), err
}
