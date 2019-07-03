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
package v1alpha1

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	scheme "github.com/pivotal/build-service-system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourceResolversGetter has a method to return a SourceResolverInterface.
// A group's client should implement this interface.
type SourceResolversGetter interface {
	SourceResolvers(namespace string) SourceResolverInterface
}

// SourceResolverInterface has methods to work with SourceResolver resources.
type SourceResolverInterface interface {
	Create(*v1alpha1.SourceResolver) (*v1alpha1.SourceResolver, error)
	Update(*v1alpha1.SourceResolver) (*v1alpha1.SourceResolver, error)
	UpdateStatus(*v1alpha1.SourceResolver) (*v1alpha1.SourceResolver, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SourceResolver, error)
	List(opts v1.ListOptions) (*v1alpha1.SourceResolverList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourceResolver, err error)
	SourceResolverExpansion
}

// sourceResolvers implements SourceResolverInterface
type sourceResolvers struct {
	client rest.Interface
	ns     string
}

// newSourceResolvers returns a SourceResolvers
func newSourceResolvers(c *BuildV1alpha1Client, namespace string) *sourceResolvers {
	return &sourceResolvers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourceResolver, and returns the corresponding sourceResolver object, and an error if there is any.
func (c *sourceResolvers) Get(name string, options v1.GetOptions) (result *v1alpha1.SourceResolver, err error) {
	result = &v1alpha1.SourceResolver{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourceResolvers that match those selectors.
func (c *sourceResolvers) List(opts v1.ListOptions) (result *v1alpha1.SourceResolverList, err error) {
	result = &v1alpha1.SourceResolverList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourceResolvers.
func (c *sourceResolvers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sourceResolver and creates it.  Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *sourceResolvers) Create(sourceResolver *v1alpha1.SourceResolver) (result *v1alpha1.SourceResolver, err error) {
	result = &v1alpha1.SourceResolver{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Body(sourceResolver).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sourceResolver and updates it. Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *sourceResolvers) Update(sourceResolver *v1alpha1.SourceResolver) (result *v1alpha1.SourceResolver, err error) {
	result = &v1alpha1.SourceResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(sourceResolver.Name).
		Body(sourceResolver).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sourceResolvers) UpdateStatus(sourceResolver *v1alpha1.SourceResolver) (result *v1alpha1.SourceResolver, err error) {
	result = &v1alpha1.SourceResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(sourceResolver.Name).
		SubResource("status").
		Body(sourceResolver).
		Do().
		Into(result)
	return
}

// Delete takes name of the sourceResolver and deletes it. Returns an error if one occurs.
func (c *sourceResolvers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourceResolvers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sourceResolver.
func (c *sourceResolvers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourceResolver, err error) {
	result = &v1alpha1.SourceResolver{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourceresolvers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
