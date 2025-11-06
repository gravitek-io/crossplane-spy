package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersionResource definitions for Crossplane resources
var (
	// Provider is a cluster-scoped resource that installs a provider package
	ProviderGVR = schema.GroupVersionResource{
		Group:    "pkg.crossplane.io",
		Version:  "v1",
		Resource: "providers",
	}

	// ProviderConfig configures a provider
	ProviderConfigGVR = schema.GroupVersionResource{
		Group:    "pkg.crossplane.io",
		Version:  "v1",
		Resource: "providerconfigs",
	}

	// CompositeResourceDefinition (XRD) defines a composite resource type
	XRDGVR = schema.GroupVersionResource{
		Group:    "apiextensions.crossplane.io",
		Version:  "v1",
		Resource: "compositeresourcedefinitions",
	}

	// Composition defines how to compose resources
	CompositionGVR = schema.GroupVersionResource{
		Group:    "apiextensions.crossplane.io",
		Version:  "v1",
		Resource: "compositions",
	}

	// Function defines a composition function
	FunctionGVR = schema.GroupVersionResource{
		Group:    "pkg.crossplane.io",
		Version:  "v1beta1",
		Resource: "functions",
	}
)

// ListProviders returns all Provider resources in the cluster
func (c *Client) ListProviders(ctx context.Context) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(ProviderGVR).List(ctx, metav1.ListOptions{})
}

// ListProviderConfigs returns all ProviderConfig resources
// Note: This is a generic method - specific provider configs may have different GVRs
func (c *Client) ListProviderConfigs(ctx context.Context, gvr schema.GroupVersionResource) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
}

// ListXRDs returns all CompositeResourceDefinition resources
func (c *Client) ListXRDs(ctx context.Context) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(XRDGVR).List(ctx, metav1.ListOptions{})
}

// ListCompositions returns all Composition resources
func (c *Client) ListCompositions(ctx context.Context) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(CompositionGVR).List(ctx, metav1.ListOptions{})
}

// ListFunctions returns all Function resources
func (c *Client) ListFunctions(ctx context.Context) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(FunctionGVR).List(ctx, metav1.ListOptions{})
}

// ListXRs returns all composite resource instances for a given XRD
// This requires the GVR to be determined from the XRD
func (c *Client) ListXRs(ctx context.Context, gvr schema.GroupVersionResource, namespace string) (*unstructured.UnstructuredList, error) {
	if namespace == "" {
		// Cluster-scoped composite resources
		return c.DynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
	}
	// Namespace-scoped composite resources
	return c.DynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{})
}

// GetResource returns a specific resource by GVR, namespace, and name
func (c *Client) GetResource(ctx context.Context, gvr schema.GroupVersionResource, namespace, name string) (*unstructured.Unstructured, error) {
	if namespace == "" {
		// Cluster-scoped resource
		return c.DynamicClient.Resource(gvr).Get(ctx, name, metav1.GetOptions{})
	}
	// Namespace-scoped resource
	return c.DynamicClient.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
}

// ListNamespaceResources returns all resources in a specific namespace
func (c *Client) ListNamespaceResources(ctx context.Context, gvr schema.GroupVersionResource, namespace string) (*unstructured.UnstructuredList, error) {
	return c.DynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{})
}
