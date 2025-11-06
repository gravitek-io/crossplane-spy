package k8s

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DiscoverXRDGVRs discovers all composite resource GVRs from XRDs
// This is used to list all composite resource instances in the cluster
func (c *Client) DiscoverXRDGVRs(ctx context.Context) ([]schema.GroupVersionResource, error) {
	xrds, err := c.ListXRDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list XRDs: %w", err)
	}

	var gvrs []schema.GroupVersionResource
	for _, xrd := range xrds.Items {
		// Get the group from spec.group
		group, found, err := getNestedString(xrd.Object, "spec", "group")
		if err != nil || !found {
			continue
		}

		// Get the plural name from spec.names.plural
		plural, found, err := getNestedString(xrd.Object, "spec", "names", "plural")
		if err != nil || !found {
			continue
		}

		// Get versions from spec.versions
		versions, found, err := getNestedSlice(xrd.Object, "spec", "versions")
		if err != nil || !found {
			continue
		}

		// Find the served and storage version
		for _, v := range versions {
			vMap, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			name, found, _ := getNestedString(vMap, "name")
			if !found {
				continue
			}

			served, _, _ := getNestedBool(vMap, "served")
			if !served {
				continue
			}

			gvr := schema.GroupVersionResource{
				Group:    group,
				Version:  name,
				Resource: plural,
			}
			gvrs = append(gvrs, gvr)

			// Typically we want the storage version, but for listing we can use any served version
			// Break after first served version for simplicity
			break
		}
	}

	return gvrs, nil
}

// DiscoverProviderConfigGVRs discovers all ProviderConfig GVRs
// ProviderConfigs have different groups depending on the provider (e.g., aws.upbound.io, gcp.upbound.io)
func (c *Client) DiscoverProviderConfigGVRs(ctx context.Context) ([]schema.GroupVersionResource, error) {
	// List all API resources
	apiGroups, err := c.Clientset.Discovery().ServerGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to discover API groups: %w", err)
	}

	var gvrs []schema.GroupVersionResource
	for _, group := range apiGroups.Groups {
		// Look for provider groups (typically contain provider name)
		// Common patterns: aws.upbound.io, gcp.upbound.io, azure.upbound.io, etc.
		groupName := group.Name
		if !isProviderGroup(groupName) {
			continue
		}

		// Get the preferred version
		version := group.PreferredVersion.Version

		// List resources in this group/version
		resourceList, err := c.Clientset.Discovery().ServerResourcesForGroupVersion(
			fmt.Sprintf("%s/%s", groupName, version),
		)
		if err != nil {
			continue
		}

		// Look for ProviderConfig resources
		for _, resource := range resourceList.APIResources {
			if strings.HasSuffix(resource.Name, "providerconfigs") || resource.Kind == "ProviderConfig" {
				gvr := schema.GroupVersionResource{
					Group:    groupName,
					Version:  version,
					Resource: resource.Name,
				}
				gvrs = append(gvrs, gvr)
			}
		}
	}

	return gvrs, nil
}

// isProviderGroup checks if an API group is likely a Crossplane provider group
func isProviderGroup(group string) bool {
	// Common provider group patterns
	providerPatterns := []string{
		".upbound.io",
		".crossplane.io",
		"aws.crossplane.io",
		"gcp.crossplane.io",
		"azure.crossplane.io",
	}

	for _, pattern := range providerPatterns {
		if strings.Contains(group, pattern) {
			return true
		}
	}

	return false
}

// IsClusterScoped checks if a resource is cluster-scoped or namespace-scoped
func (c *Client) IsClusterScoped(ctx context.Context, gvr schema.GroupVersionResource) (bool, error) {
	resourceList, err := c.Clientset.Discovery().ServerResourcesForGroupVersion(
		fmt.Sprintf("%s/%s", gvr.Group, gvr.Version),
	)
	if err != nil {
		return false, fmt.Errorf("failed to discover resources: %w", err)
	}

	for _, resource := range resourceList.APIResources {
		if resource.Name == gvr.Resource {
			return !resource.Namespaced, nil
		}
	}

	return false, fmt.Errorf("resource not found: %s", gvr.Resource)
}

// Helper functions to extract nested fields from unstructured objects

func getNestedString(obj map[string]interface{}, fields ...string) (string, bool, error) {
	val, found, err := getNestedField(obj, fields...)
	if err != nil || !found {
		return "", found, err
	}
	str, ok := val.(string)
	if !ok {
		return "", false, fmt.Errorf("value is not a string")
	}
	return str, true, nil
}

func getNestedBool(obj map[string]interface{}, fields ...string) (bool, bool, error) {
	val, found, err := getNestedField(obj, fields...)
	if err != nil || !found {
		return false, found, err
	}
	b, ok := val.(bool)
	if !ok {
		return false, false, fmt.Errorf("value is not a bool")
	}
	return b, true, nil
}

func getNestedSlice(obj map[string]interface{}, fields ...string) ([]interface{}, bool, error) {
	val, found, err := getNestedField(obj, fields...)
	if err != nil || !found {
		return nil, found, err
	}
	slice, ok := val.([]interface{})
	if !ok {
		return nil, false, fmt.Errorf("value is not a slice")
	}
	return slice, true, nil
}

func getNestedField(obj map[string]interface{}, fields ...string) (interface{}, bool, error) {
	var val interface{} = obj

	for _, field := range fields {
		m, ok := val.(map[string]interface{})
		if !ok {
			return nil, false, fmt.Errorf("value is not a map")
		}
		val, ok = m[field]
		if !ok {
			return nil, false, nil
		}
	}

	return val, true, nil
}
