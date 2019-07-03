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
package fake

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSourceResolvers implements SourceResolverInterface
type FakeSourceResolvers struct {
	Fake *FakeBuildV1alpha1
	ns   string
}

var sourceresolversResource = schema.GroupVersionResource{Group: "build.pivotal.io", Version: "v1alpha1", Resource: "sourceresolvers"}

var sourceresolversKind = schema.GroupVersionKind{Group: "build.pivotal.io", Version: "v1alpha1", Kind: "SourceResolver"}

// Get takes name of the sourceResolver, and returns the corresponding sourceResolver object, and an error if there is any.
func (c *FakeSourceResolvers) Get(name string, options v1.GetOptions) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sourceresolversResource, c.ns, name), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// List takes label and field selectors, and returns the list of SourceResolvers that match those selectors.
func (c *FakeSourceResolvers) List(opts v1.ListOptions) (result *v1alpha1.SourceResolverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sourceresolversResource, sourceresolversKind, c.ns, opts), &v1alpha1.SourceResolverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SourceResolverList{ListMeta: obj.(*v1alpha1.SourceResolverList).ListMeta}
	for _, item := range obj.(*v1alpha1.SourceResolverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sourceResolvers.
func (c *FakeSourceResolvers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sourceresolversResource, c.ns, opts))

}

// Create takes the representation of a sourceResolver and creates it.  Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *FakeSourceResolvers) Create(sourceResolver *v1alpha1.SourceResolver) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sourceresolversResource, c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// Update takes the representation of a sourceResolver and updates it. Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *FakeSourceResolvers) Update(sourceResolver *v1alpha1.SourceResolver) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sourceresolversResource, c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSourceResolvers) UpdateStatus(sourceResolver *v1alpha1.SourceResolver) (*v1alpha1.SourceResolver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sourceresolversResource, "status", c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// Delete takes name of the sourceResolver and deletes it. Returns an error if one occurs.
func (c *FakeSourceResolvers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sourceresolversResource, c.ns, name), &v1alpha1.SourceResolver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSourceResolvers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sourceresolversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SourceResolverList{})
	return err
}

// Patch applies the patch and returns the patched sourceResolver.
func (c *FakeSourceResolvers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sourceresolversResource, c.ns, name, data, subresources...), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}
