/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/xianfuzhang/baremetal-controller/pkg/apis/baremetal/v1alpha1"
	scheme "github.com/xianfuzhang/baremetal-controller/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BaremetalsGetter has a method to return a BaremetalInterface.
// A group's client should implement this interface.
type BaremetalsGetter interface {
	Baremetals(namespace string) BaremetalInterface
}

// BaremetalInterface has methods to work with Baremetal resources.
type BaremetalInterface interface {
	Create(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.CreateOptions) (*v1alpha1.Baremetal, error)
	Update(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (*v1alpha1.Baremetal, error)
	UpdateStatus(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (*v1alpha1.Baremetal, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Baremetal, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BaremetalList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Baremetal, err error)
	BaremetalExpansion
}

// baremetals implements BaremetalInterface
type baremetals struct {
	client rest.Interface
	ns     string
}

// newBaremetals returns a Baremetals
func newBaremetals(c *GroupV1alpha1Client, namespace string) *baremetals {
	return &baremetals{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the baremetal, and returns the corresponding baremetal object, and an error if there is any.
func (c *baremetals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Baremetal, err error) {
	result = &v1alpha1.Baremetal{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("baremetals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Baremetals that match those selectors.
func (c *baremetals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BaremetalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BaremetalList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("baremetals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested baremetals.
func (c *baremetals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("baremetals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a baremetal and creates it.  Returns the server's representation of the baremetal, and an error, if there is any.
func (c *baremetals) Create(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.CreateOptions) (result *v1alpha1.Baremetal, err error) {
	result = &v1alpha1.Baremetal{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("baremetals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(baremetal).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a baremetal and updates it. Returns the server's representation of the baremetal, and an error, if there is any.
func (c *baremetals) Update(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (result *v1alpha1.Baremetal, err error) {
	result = &v1alpha1.Baremetal{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("baremetals").
		Name(baremetal.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(baremetal).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *baremetals) UpdateStatus(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (result *v1alpha1.Baremetal, err error) {
	result = &v1alpha1.Baremetal{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("baremetals").
		Name(baremetal.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(baremetal).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the baremetal and deletes it. Returns an error if one occurs.
func (c *baremetals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("baremetals").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *baremetals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("baremetals").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched baremetal.
func (c *baremetals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Baremetal, err error) {
	result = &v1alpha1.Baremetal{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("baremetals").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}