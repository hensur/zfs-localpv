/*
Copyright 2019 The OpenEBS Authors

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

	v1 "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZFSNodes implements ZFSNodeInterface
type FakeZFSNodes struct {
	Fake *FakeZfsV1
	ns   string
}

var zfsnodesResource = v1.SchemeGroupVersion.WithResource("zfsnodes")

var zfsnodesKind = v1.SchemeGroupVersion.WithKind("ZFSNode")

// Get takes name of the zFSNode, and returns the corresponding zFSNode object, and an error if there is any.
func (c *FakeZFSNodes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ZFSNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zfsnodesResource, c.ns, name), &v1.ZFSNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ZFSNode), err
}

// List takes label and field selectors, and returns the list of ZFSNodes that match those selectors.
func (c *FakeZFSNodes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ZFSNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zfsnodesResource, zfsnodesKind, c.ns, opts), &v1.ZFSNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ZFSNodeList{ListMeta: obj.(*v1.ZFSNodeList).ListMeta}
	for _, item := range obj.(*v1.ZFSNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zFSNodes.
func (c *FakeZFSNodes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zfsnodesResource, c.ns, opts))

}

// Create takes the representation of a zFSNode and creates it.  Returns the server's representation of the zFSNode, and an error, if there is any.
func (c *FakeZFSNodes) Create(ctx context.Context, zFSNode *v1.ZFSNode, opts metav1.CreateOptions) (result *v1.ZFSNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zfsnodesResource, c.ns, zFSNode), &v1.ZFSNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ZFSNode), err
}

// Update takes the representation of a zFSNode and updates it. Returns the server's representation of the zFSNode, and an error, if there is any.
func (c *FakeZFSNodes) Update(ctx context.Context, zFSNode *v1.ZFSNode, opts metav1.UpdateOptions) (result *v1.ZFSNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zfsnodesResource, c.ns, zFSNode), &v1.ZFSNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ZFSNode), err
}

// Delete takes name of the zFSNode and deletes it. Returns an error if one occurs.
func (c *FakeZFSNodes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(zfsnodesResource, c.ns, name, opts), &v1.ZFSNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZFSNodes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zfsnodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ZFSNodeList{})
	return err
}

// Patch applies the patch and returns the patched zFSNode.
func (c *FakeZFSNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ZFSNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zfsnodesResource, c.ns, name, pt, data, subresources...), &v1.ZFSNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ZFSNode), err
}
