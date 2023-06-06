// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/networking/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MultiClusterServicesGetter has a method to return a MultiClusterServiceInterface.
// A group's client should implement this interface.
type MultiClusterServicesGetter interface {
	MultiClusterServices(namespace string) MultiClusterServiceInterface
}

// MultiClusterServiceInterface has methods to work with MultiClusterService resources.
type MultiClusterServiceInterface interface {
	Create(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.CreateOptions) (*v1alpha1.MultiClusterService, error)
	Update(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (*v1alpha1.MultiClusterService, error)
	UpdateStatus(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (*v1alpha1.MultiClusterService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MultiClusterService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MultiClusterServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterService, err error)
	MultiClusterServiceExpansion
}

// multiClusterServices implements MultiClusterServiceInterface
type multiClusterServices struct {
	client rest.Interface
	ns     string
}

// newMultiClusterServices returns a MultiClusterServices
func newMultiClusterServices(c *NetworkingV1alpha1Client, namespace string) *multiClusterServices {
	return &multiClusterServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multiClusterService, and returns the corresponding multiClusterService object, and an error if there is any.
func (c *multiClusterServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterService, err error) {
	result = &v1alpha1.MultiClusterService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultiClusterServices that match those selectors.
func (c *multiClusterServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MultiClusterServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multiClusterServices.
func (c *multiClusterServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multiClusterService and creates it.  Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *multiClusterServices) Create(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.CreateOptions) (result *v1alpha1.MultiClusterService, err error) {
	result = &v1alpha1.MultiClusterService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multiclusterservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multiClusterService and updates it. Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *multiClusterServices) Update(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterService, err error) {
	result = &v1alpha1.MultiClusterService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterservices").
		Name(multiClusterService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *multiClusterServices) UpdateStatus(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterService, err error) {
	result = &v1alpha1.MultiClusterService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterservices").
		Name(multiClusterService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multiClusterService and deletes it. Returns an error if one occurs.
func (c *multiClusterServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multiClusterServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multiClusterService.
func (c *multiClusterServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterService, err error) {
	result = &v1alpha1.MultiClusterService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multiclusterservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
