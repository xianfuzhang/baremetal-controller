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

package fake

import (
	"context"

	v1alpha1 "github.com/xianfuzhang/baremetal-controller/pkg/apis/baremetal/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBaremetals implements BaremetalInterface
type FakeBaremetals struct {
	Fake *FakeGroupV1alpha1
	ns   string
}

var baremetalsResource = schema.GroupVersionResource{Group: "group.nocsys.cn", Version: "v1alpha1", Resource: "baremetals"}

var baremetalsKind = schema.GroupVersionKind{Group: "group.nocsys.cn", Version: "v1alpha1", Kind: "Baremetal"}

// Get takes name of the baremetal, and returns the corresponding baremetal object, and an error if there is any.
func (c *FakeBaremetals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Baremetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(baremetalsResource, c.ns, name), &v1alpha1.Baremetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Baremetal), err
}

// List takes label and field selectors, and returns the list of Baremetals that match those selectors.
func (c *FakeBaremetals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BaremetalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(baremetalsResource, baremetalsKind, c.ns, opts), &v1alpha1.BaremetalList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BaremetalList{ListMeta: obj.(*v1alpha1.BaremetalList).ListMeta}
	for _, item := range obj.(*v1alpha1.BaremetalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested baremetals.
func (c *FakeBaremetals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(baremetalsResource, c.ns, opts))

}

// Create takes the representation of a baremetal and creates it.  Returns the server's representation of the baremetal, and an error, if there is any.
func (c *FakeBaremetals) Create(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.CreateOptions) (result *v1alpha1.Baremetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(baremetalsResource, c.ns, baremetal), &v1alpha1.Baremetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Baremetal), err
}

// Update takes the representation of a baremetal and updates it. Returns the server's representation of the baremetal, and an error, if there is any.
func (c *FakeBaremetals) Update(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (result *v1alpha1.Baremetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(baremetalsResource, c.ns, baremetal), &v1alpha1.Baremetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Baremetal), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBaremetals) UpdateStatus(ctx context.Context, baremetal *v1alpha1.Baremetal, opts v1.UpdateOptions) (*v1alpha1.Baremetal, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(baremetalsResource, "status", c.ns, baremetal), &v1alpha1.Baremetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Baremetal), err
}

// Delete takes name of the baremetal and deletes it. Returns an error if one occurs.
func (c *FakeBaremetals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(baremetalsResource, c.ns, name), &v1alpha1.Baremetal{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBaremetals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(baremetalsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BaremetalList{})
	return err
}

// Patch applies the patch and returns the patched baremetal.
func (c *FakeBaremetals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Baremetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(baremetalsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Baremetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Baremetal), err
}
