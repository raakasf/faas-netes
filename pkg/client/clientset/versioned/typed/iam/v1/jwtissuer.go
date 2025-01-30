/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openfaas/faas-netes/pkg/apis/iam/v1"
	scheme "github.com/openfaas/faas-netes/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JwtIssuersGetter has a method to return a JwtIssuerInterface.
// A group's client should implement this interface.
type JwtIssuersGetter interface {
	JwtIssuers(namespace string) JwtIssuerInterface
}

// JwtIssuerInterface has methods to work with JwtIssuer resources.
type JwtIssuerInterface interface {
	Create(ctx context.Context, jwtIssuer *v1.JwtIssuer, opts metav1.CreateOptions) (*v1.JwtIssuer, error)
	Update(ctx context.Context, jwtIssuer *v1.JwtIssuer, opts metav1.UpdateOptions) (*v1.JwtIssuer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.JwtIssuer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.JwtIssuerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.JwtIssuer, err error)
	JwtIssuerExpansion
}

// jwtIssuers implements JwtIssuerInterface
type jwtIssuers struct {
	client rest.Interface
	ns     string
}

// newJwtIssuers returns a JwtIssuers
func newJwtIssuers(c *IamV1Client, namespace string) *jwtIssuers {
	return &jwtIssuers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jwtIssuer, and returns the corresponding jwtIssuer object, and an error if there is any.
func (c *jwtIssuers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.JwtIssuer, err error) {
	result = &v1.JwtIssuer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jwtissuers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JwtIssuers that match those selectors.
func (c *jwtIssuers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.JwtIssuerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.JwtIssuerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jwtissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jwtIssuers.
func (c *jwtIssuers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jwtissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jwtIssuer and creates it.  Returns the server's representation of the jwtIssuer, and an error, if there is any.
func (c *jwtIssuers) Create(ctx context.Context, jwtIssuer *v1.JwtIssuer, opts metav1.CreateOptions) (result *v1.JwtIssuer, err error) {
	result = &v1.JwtIssuer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jwtissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jwtIssuer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jwtIssuer and updates it. Returns the server's representation of the jwtIssuer, and an error, if there is any.
func (c *jwtIssuers) Update(ctx context.Context, jwtIssuer *v1.JwtIssuer, opts metav1.UpdateOptions) (result *v1.JwtIssuer, err error) {
	result = &v1.JwtIssuer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jwtissuers").
		Name(jwtIssuer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jwtIssuer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jwtIssuer and deletes it. Returns an error if one occurs.
func (c *jwtIssuers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jwtissuers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jwtIssuers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jwtissuers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jwtIssuer.
func (c *jwtIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.JwtIssuer, err error) {
	result = &v1.JwtIssuer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jwtissuers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
