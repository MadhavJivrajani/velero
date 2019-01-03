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

// FakeVolumeSnapshotLocations implements VolumeSnapshotLocationInterface
type FakeVolumeSnapshotLocations struct {
	Fake *FakeArkV1
	ns   string
}

var volumesnapshotlocationsResource = schema.GroupVersionResource{Group: "ark.heptio.com", Version: "v1", Resource: "volumesnapshotlocations"}

var volumesnapshotlocationsKind = schema.GroupVersionKind{Group: "ark.heptio.com", Version: "v1", Kind: "VolumeSnapshotLocation"}

// Get takes name of the volumeSnapshotLocation, and returns the corresponding volumeSnapshotLocation object, and an error if there is any.
func (c *FakeVolumeSnapshotLocations) Get(name string, options v1.GetOptions) (result *ark_v1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotlocationsResource, c.ns, name), &ark_v1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.VolumeSnapshotLocation), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotLocations that match those selectors.
func (c *FakeVolumeSnapshotLocations) List(opts v1.ListOptions) (result *ark_v1.VolumeSnapshotLocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotlocationsResource, volumesnapshotlocationsKind, c.ns, opts), &ark_v1.VolumeSnapshotLocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ark_v1.VolumeSnapshotLocationList{ListMeta: obj.(*ark_v1.VolumeSnapshotLocationList).ListMeta}
	for _, item := range obj.(*ark_v1.VolumeSnapshotLocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotLocations.
func (c *FakeVolumeSnapshotLocations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotlocationsResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshotLocation and creates it.  Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *FakeVolumeSnapshotLocations) Create(volumeSnapshotLocation *ark_v1.VolumeSnapshotLocation) (result *ark_v1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotlocationsResource, c.ns, volumeSnapshotLocation), &ark_v1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.VolumeSnapshotLocation), err
}

// Update takes the representation of a volumeSnapshotLocation and updates it. Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *FakeVolumeSnapshotLocations) Update(volumeSnapshotLocation *ark_v1.VolumeSnapshotLocation) (result *ark_v1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotlocationsResource, c.ns, volumeSnapshotLocation), &ark_v1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.VolumeSnapshotLocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshotLocations) UpdateStatus(volumeSnapshotLocation *ark_v1.VolumeSnapshotLocation) (*ark_v1.VolumeSnapshotLocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotlocationsResource, "status", c.ns, volumeSnapshotLocation), &ark_v1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.VolumeSnapshotLocation), err
}

// Delete takes name of the volumeSnapshotLocation and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotLocations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesnapshotlocationsResource, c.ns, name), &ark_v1.VolumeSnapshotLocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotLocations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotlocationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ark_v1.VolumeSnapshotLocationList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotLocation.
func (c *FakeVolumeSnapshotLocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ark_v1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotlocationsResource, c.ns, name, data, subresources...), &ark_v1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ark_v1.VolumeSnapshotLocation), err
}
