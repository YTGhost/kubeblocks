/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	scheme "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterDefinitionsGetter has a method to return a ClusterDefinitionInterface.
// A group's client should implement this interface.
type ClusterDefinitionsGetter interface {
	ClusterDefinitions() ClusterDefinitionInterface
}

// ClusterDefinitionInterface has methods to work with ClusterDefinition resources.
type ClusterDefinitionInterface interface {
	Create(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.CreateOptions) (*v1alpha1.ClusterDefinition, error)
	Update(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (*v1alpha1.ClusterDefinition, error)
	UpdateStatus(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (*v1alpha1.ClusterDefinition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterDefinition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterDefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterDefinition, err error)
	ClusterDefinitionExpansion
}

// clusterDefinitions implements ClusterDefinitionInterface
type clusterDefinitions struct {
	client rest.Interface
}

// newClusterDefinitions returns a ClusterDefinitions
func newClusterDefinitions(c *AppsV1alpha1Client) *clusterDefinitions {
	return &clusterDefinitions{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterDefinition, and returns the corresponding clusterDefinition object, and an error if there is any.
func (c *clusterDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterDefinition, err error) {
	result = &v1alpha1.ClusterDefinition{}
	err = c.client.Get().
		Resource("clusterdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterDefinitions that match those selectors.
func (c *clusterDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterDefinitionList{}
	err = c.client.Get().
		Resource("clusterdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterDefinitions.
func (c *clusterDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterDefinition and creates it.  Returns the server's representation of the clusterDefinition, and an error, if there is any.
func (c *clusterDefinitions) Create(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.CreateOptions) (result *v1alpha1.ClusterDefinition, err error) {
	result = &v1alpha1.ClusterDefinition{}
	err = c.client.Post().
		Resource("clusterdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterDefinition and updates it. Returns the server's representation of the clusterDefinition, and an error, if there is any.
func (c *clusterDefinitions) Update(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (result *v1alpha1.ClusterDefinition, err error) {
	result = &v1alpha1.ClusterDefinition{}
	err = c.client.Put().
		Resource("clusterdefinitions").
		Name(clusterDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDefinition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterDefinitions) UpdateStatus(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (result *v1alpha1.ClusterDefinition, err error) {
	result = &v1alpha1.ClusterDefinition{}
	err = c.client.Put().
		Resource("clusterdefinitions").
		Name(clusterDefinition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterDefinition and deletes it. Returns an error if one occurs.
func (c *clusterDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterdefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterdefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterDefinition.
func (c *clusterDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterDefinition, err error) {
	result = &v1alpha1.ClusterDefinition{}
	err = c.client.Patch(pt).
		Resource("clusterdefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
