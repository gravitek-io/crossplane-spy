package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitek/crossplane-spy/internal/k8s"
	"github.com/gravitek/crossplane-spy/internal/models"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// getResources returns all Crossplane resources summary
func getResources(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		// Get all resource types
		providers, _ := client.ListProviders(ctx)
		xrds, _ := client.ListXRDs(ctx)
		compositions, _ := client.ListCompositions(ctx)
		functions, _ := client.ListFunctions(ctx)

		summary := models.ResourceSummary{
			Providers:    len(providers.Items),
			XRDs:         len(xrds.Items),
			Compositions: len(compositions.Items),
			Functions:    len(functions.Items),
		}

		c.JSON(http.StatusOK, summary)
	}
}

// getResourcesByKind returns resources of a specific kind
func getResourcesByKind(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		kind := c.Param("kind")
		c.JSON(http.StatusOK, gin.H{
			"message": "List resources by kind",
			"kind":    kind,
		})
	}
}

// getResource returns a specific resource
func getResource(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		kind := c.Param("kind")
		namespace := c.Param("namespace")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"kind":      kind,
			"namespace": namespace,
			"name":      name,
		})
	}
}

// getProviders returns all Provider resources
func getProviders(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		providerList, err := client.ListProviders(ctx)
		if err != nil {
			log.Printf("Error listing providers: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list providers"})
			return
		}

		providers := convertToProviders(providerList.Items)
		c.JSON(http.StatusOK, gin.H{
			"kind":  "ProviderList",
			"count": len(providers),
			"items": providers,
		})
	}
}

// getProviderConfigs returns all ProviderConfig resources
func getProviderConfigs(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		gvrs, err := client.DiscoverProviderConfigGVRs(ctx)
		if err != nil {
			log.Printf("Error discovering provider configs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to discover provider configs"})
			return
		}

		var allConfigs []models.ProviderConfig
		for _, gvr := range gvrs {
			configs, err := client.ListProviderConfigs(ctx, gvr)
			if err != nil {
				log.Printf("Error listing provider configs for %v: %v", gvr, err)
				continue
			}
			allConfigs = append(allConfigs, convertToProviderConfigs(configs.Items)...)
		}

		c.JSON(http.StatusOK, gin.H{
			"kind":  "ProviderConfigList",
			"count": len(allConfigs),
			"items": allConfigs,
		})
	}
}

// getXRDs returns all XRD resources
func getXRDs(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		xrdList, err := client.ListXRDs(ctx)
		if err != nil {
			log.Printf("Error listing XRDs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list XRDs"})
			return
		}

		xrds := convertToXRDs(xrdList.Items)
		c.JSON(http.StatusOK, gin.H{
			"kind":  "CompositeResourceDefinitionList",
			"count": len(xrds),
			"items": xrds,
		})
	}
}

// getCompositions returns all Composition resources
func getCompositions(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		compList, err := client.ListCompositions(ctx)
		if err != nil {
			log.Printf("Error listing compositions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list compositions"})
			return
		}

		compositions := convertToCompositions(compList.Items)
		c.JSON(http.StatusOK, gin.H{
			"kind":  "CompositionList",
			"count": len(compositions),
			"items": compositions,
		})
	}
}

// getXRs returns all Composite Resource (XR) instances
func getXRs(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		gvrs, err := client.DiscoverXRDGVRs(ctx)
		if err != nil {
			log.Printf("Error discovering XRs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to discover composite resources"})
			return
		}

		var allXRs []models.CompositeResource
		for _, gvr := range gvrs {
			xrs, err := client.ListXRs(ctx, gvr, "")
			if err != nil {
				log.Printf("Error listing XRs for %v: %v", gvr, err)
				continue
			}
			allXRs = append(allXRs, convertToCompositeResources(xrs.Items)...)
		}

		c.JSON(http.StatusOK, gin.H{
			"kind":  "CompositeResourceList",
			"count": len(allXRs),
			"items": allXRs,
		})
	}
}

// getFunctions returns all Function resources
func getFunctions(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		funcList, err := client.ListFunctions(ctx)
		if err != nil {
			log.Printf("Error listing functions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list functions"})
			return
		}

		functions := convertToFunctions(funcList.Items)
		c.JSON(http.StatusOK, gin.H{
			"kind":  "FunctionList",
			"count": len(functions),
			"items": functions,
		})
	}
}

// getClusterResources returns cluster-scoped Crossplane resources
func getClusterResources(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		providers, _ := client.ListProviders(ctx)
		xrds, _ := client.ListXRDs(ctx)
		compositions, _ := client.ListCompositions(ctx)
		functions, _ := client.ListFunctions(ctx)

		var resources []interface{}
		resources = append(resources, convertToAnySlice(convertToProviders(providers.Items))...)
		resources = append(resources, convertToAnySlice(convertToXRDs(xrds.Items))...)
		resources = append(resources, convertToAnySlice(convertToCompositions(compositions.Items))...)
		resources = append(resources, convertToAnySlice(convertToFunctions(functions.Items))...)

		c.JSON(http.StatusOK, gin.H{
			"scope": "cluster",
			"count": len(resources),
			"items": resources,
		})
	}
}

// getNamespaceResources returns namespace-scoped Crossplane resources
func getNamespaceResources(client *k8s.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		gvrs, err := client.DiscoverXRDGVRs(ctx)
		if err != nil {
			log.Printf("Error discovering namespace resources: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to discover namespace resources"})
			return
		}

		var resources []interface{}
		for _, gvr := range gvrs {
			// Try to list with namespace (will fail for cluster-scoped)
			namespaces := []string{"default", "crossplane-system"}
			for _, ns := range namespaces {
				xrs, err := client.ListXRs(ctx, gvr, ns)
				if err != nil {
					continue
				}
				resources = append(resources, convertToAnySlice(convertToCompositeResources(xrs.Items))...)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"scope": "namespace",
			"count": len(resources),
			"items": resources,
		})
	}
}

// Converter functions

func convertToProviders(items []unstructured.Unstructured) []models.Provider {
	providers := make([]models.Provider, 0, len(items))
	for _, item := range items {
		provider := models.Provider{
			BaseResource: models.ConvertToBaseResource(&item, models.ScopeCluster),
			Status:       models.ProviderStatus{ResourceStatus: models.ConvertToResourceStatus(&item)},
		}
		providers = append(providers, provider)
	}
	return providers
}

func convertToProviderConfigs(items []unstructured.Unstructured) []models.ProviderConfig {
	configs := make([]models.ProviderConfig, 0, len(items))
	for _, item := range items {
		config := models.ProviderConfig{
			BaseResource: models.ConvertToBaseResource(&item, models.ScopeCluster),
			Status:       models.ConvertToResourceStatus(&item),
		}
		configs = append(configs, config)
	}
	return configs
}

func convertToXRDs(items []unstructured.Unstructured) []models.XRD {
	xrds := make([]models.XRD, 0, len(items))
	for _, item := range items {
		xrd := models.XRD{
			BaseResource: models.ConvertToBaseResource(&item, models.ScopeCluster),
			Status:       models.XRDStatus{ResourceStatus: models.ConvertToResourceStatus(&item)},
		}
		xrds = append(xrds, xrd)
	}
	return xrds
}

func convertToCompositions(items []unstructured.Unstructured) []models.Composition {
	compositions := make([]models.Composition, 0, len(items))
	for _, item := range items {
		composition := models.Composition{
			BaseResource: models.ConvertToBaseResource(&item, models.ScopeCluster),
			Status:       models.CompositionStatus{ResourceStatus: models.ConvertToResourceStatus(&item)},
		}
		compositions = append(compositions, composition)
	}
	return compositions
}

func convertToFunctions(items []unstructured.Unstructured) []models.Function {
	functions := make([]models.Function, 0, len(items))
	for _, item := range items {
		function := models.Function{
			BaseResource: models.ConvertToBaseResource(&item, models.ScopeCluster),
			Status:       models.FunctionStatus{ResourceStatus: models.ConvertToResourceStatus(&item)},
		}
		functions = append(functions, function)
	}
	return functions
}

func convertToCompositeResources(items []unstructured.Unstructured) []models.CompositeResource {
	xrs := make([]models.CompositeResource, 0, len(items))
	for _, item := range items {
		scope := models.ScopeCluster
		if item.GetNamespace() != "" {
			scope = models.ScopeNamespace
		}
		xr := models.CompositeResource{
			BaseResource: models.ConvertToBaseResource(&item, scope),
			Status:       models.CompositeResourceStatus{ResourceStatus: models.ConvertToResourceStatus(&item)},
		}
		xrs = append(xrs, xr)
	}
	return xrs
}

// Generic converter to []interface{}
func convertToAnySlice[T any](items []T) []interface{} {
	result := make([]interface{}, len(items))
	for i, item := range items {
		result[i] = item
	}
	return result
}
