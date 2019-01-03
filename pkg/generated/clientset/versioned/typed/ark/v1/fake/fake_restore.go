/*
Copyright the Heptio Ark contributors.

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
	ark_v1 "github.com/heptio/ark/pkg/apis/ark/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRestores implements RestoreInterface
type FakeRestores struct {
	Fake *FakeArkV1
	ns   string
}

var restoresResource = schema.GroupVersionResource{Group: "ark.heptio.com", Version: "v1", Resource: "restores"}

var restoresKind = schema.GroupVersionKind{Group: "ark.heptio.com", Version: "v1", Kind: "Restore"}

// Get takes name of the restore, and returns the corresponding restore object, and an error if there is any.
func (c *FakeRestores) Get(name string, options v1.GetOptions) (result *ark_v1.Restore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(restoresResource, c.ns, name), &ark_v1.Restore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.Restore), err
}

// List takes label and field selectors, and returns the list of Restores that match those selectors.
func (c *FakeRestores) List(opts v1.ListOptions) (result *ark_v1.RestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(restoresResource, restoresKind, c.ns, opts), &ark_v1.RestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ark_v1.RestoreList{ListMeta: obj.(*ark_v1.RestoreList).ListMeta}
	for _, item := range obj.(*ark_v1.RestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested restores.
func (c *FakeRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(restoresResource, c.ns, opts))

}

// Create takes the representation of a restore and creates it.  Returns the server's representation of the restore, and an error, if there is any.
func (c *FakeRestores) Create(restore *ark_v1.Restore) (result *ark_v1.Restore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(restoresResource, c.ns, restore), &ark_v1.Restore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.Restore), err
}

// Update takes the representation of a restore and updates it. Returns the server's representation of the restore, and an error, if there is any.
func (c *FakeRestores) Update(restore *ark_v1.Restore) (result *ark_v1.Restore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(restoresResource, c.ns, restore), &ark_v1.Restore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.Restore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRestores) UpdateStatus(restore *ark_v1.Restore) (*ark_v1.Restore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(restoresResource, "status", c.ns, restore), &ark_v1.Restore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.Restore), err
}

// Delete takes name of the restore and deletes it. Returns an error if one occurs.
func (c *FakeRestores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(restoresResource, c.ns, name), &ark_v1.Restore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(restoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ark_v1.RestoreList{})
	return err
}

// Patch applies the patch and returns the patched restore.
func (c *FakeRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ark_v1.Restore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(restoresResource, c.ns, name, data, subresources...), &ark_v1.Restore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.Restore), err
}
