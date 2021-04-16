/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterStacks implements ClusterStackInterface
type FakeClusterStacks struct {
	Fake *FakeKpackV1alpha1
}

var clusterstacksResource = schema.GroupVersionResource{Group: "kpack.io", Version: "v1alpha1", Resource: "clusterstacks"}

var clusterstacksKind = schema.GroupVersionKind{Group: "kpack.io", Version: "v1alpha1", Kind: "ClusterStack"}

// Get takes name of the clusterStack, and returns the corresponding clusterStack object, and an error if there is any.
func (c *FakeClusterStacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterstacksResource, name), &v1alpha1.ClusterStack{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStack), err
}

// List takes label and field selectors, and returns the list of ClusterStacks that match those selectors.
func (c *FakeClusterStacks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterStackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterstacksResource, clusterstacksKind, opts), &v1alpha1.ClusterStackList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterStackList{ListMeta: obj.(*v1alpha1.ClusterStackList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterStackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterStacks.
func (c *FakeClusterStacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterstacksResource, opts))
}

// Create takes the representation of a clusterStack and creates it.  Returns the server's representation of the clusterStack, and an error, if there is any.
func (c *FakeClusterStacks) Create(ctx context.Context, clusterStack *v1alpha1.ClusterStack, opts v1.CreateOptions) (result *v1alpha1.ClusterStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterstacksResource, clusterStack), &v1alpha1.ClusterStack{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStack), err
}

// Update takes the representation of a clusterStack and updates it. Returns the server's representation of the clusterStack, and an error, if there is any.
func (c *FakeClusterStacks) Update(ctx context.Context, clusterStack *v1alpha1.ClusterStack, opts v1.UpdateOptions) (result *v1alpha1.ClusterStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterstacksResource, clusterStack), &v1alpha1.ClusterStack{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterStacks) UpdateStatus(ctx context.Context, clusterStack *v1alpha1.ClusterStack, opts v1.UpdateOptions) (*v1alpha1.ClusterStack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterstacksResource, "status", clusterStack), &v1alpha1.ClusterStack{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStack), err
}

// Delete takes name of the clusterStack and deletes it. Returns an error if one occurs.
func (c *FakeClusterStacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterstacksResource, name), &v1alpha1.ClusterStack{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterStacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterstacksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterStackList{})
	return err
}

// Patch applies the patch and returns the patched clusterStack.
func (c *FakeClusterStacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterstacksResource, name, pt, data, subresources...), &v1alpha1.ClusterStack{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStack), err
}
