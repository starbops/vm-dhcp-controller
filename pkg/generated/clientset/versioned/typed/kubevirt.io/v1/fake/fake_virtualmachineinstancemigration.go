/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "kubevirt.io/api/core/v1"
)

// FakeVirtualMachineInstanceMigrations implements VirtualMachineInstanceMigrationInterface
type FakeVirtualMachineInstanceMigrations struct {
	Fake *FakeKubevirtV1
	ns   string
}

var virtualmachineinstancemigrationsResource = v1.SchemeGroupVersion.WithResource("virtualmachineinstancemigrations")

var virtualmachineinstancemigrationsKind = v1.SchemeGroupVersion.WithKind("VirtualMachineInstanceMigration")

// Get takes name of the virtualMachineInstanceMigration, and returns the corresponding virtualMachineInstanceMigration object, and an error if there is any.
func (c *FakeVirtualMachineInstanceMigrations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancemigrationsResource, c.ns, name), &v1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineInstanceMigration), err
}

// List takes label and field selectors, and returns the list of VirtualMachineInstanceMigrations that match those selectors.
func (c *FakeVirtualMachineInstanceMigrations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstanceMigrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancemigrationsResource, virtualmachineinstancemigrationsKind, c.ns, opts), &v1.VirtualMachineInstanceMigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VirtualMachineInstanceMigrationList{ListMeta: obj.(*v1.VirtualMachineInstanceMigrationList).ListMeta}
	for _, item := range obj.(*v1.VirtualMachineInstanceMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstanceMigrations.
func (c *FakeVirtualMachineInstanceMigrations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancemigrationsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineInstanceMigration and creates it.  Returns the server's representation of the virtualMachineInstanceMigration, and an error, if there is any.
func (c *FakeVirtualMachineInstanceMigrations) Create(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.CreateOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancemigrationsResource, c.ns, virtualMachineInstanceMigration), &v1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineInstanceMigration), err
}

// Update takes the representation of a virtualMachineInstanceMigration and updates it. Returns the server's representation of the virtualMachineInstanceMigration, and an error, if there is any.
func (c *FakeVirtualMachineInstanceMigrations) Update(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancemigrationsResource, c.ns, virtualMachineInstanceMigration), &v1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineInstanceMigration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineInstanceMigrations) UpdateStatus(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceMigration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancemigrationsResource, "status", c.ns, virtualMachineInstanceMigration), &v1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineInstanceMigration), err
}

// Delete takes name of the virtualMachineInstanceMigration and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineInstanceMigrations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualmachineinstancemigrationsResource, c.ns, name, opts), &v1.VirtualMachineInstanceMigration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineInstanceMigrations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineinstancemigrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.VirtualMachineInstanceMigrationList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineInstanceMigration.
func (c *FakeVirtualMachineInstanceMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancemigrationsResource, c.ns, name, pt, data, subresources...), &v1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineInstanceMigration), err
}