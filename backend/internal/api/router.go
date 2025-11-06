package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravitek/crossplane-spy/internal/k8s"
)

// NewRouter creates and configures the API router
func NewRouter(k8sClient *k8s.Client) *gin.Engine {
	router := gin.Default()

	// CORS middleware for Next.js frontend
	router.Use(corsMiddleware())

	// Health check endpoint
	router.GET("/health", healthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Resource endpoints
		v1.GET("/resources", getResources(k8sClient))
		v1.GET("/resources/:kind", getResourcesByKind(k8sClient))
		v1.GET("/resources/:kind/:namespace/:name", getResource(k8sClient))

		// Specific resource type endpoints
		v1.GET("/providers", getProviders(k8sClient))
		v1.GET("/providerconfigs", getProviderConfigs(k8sClient))
		v1.GET("/xrds", getXRDs(k8sClient))
		v1.GET("/compositions", getCompositions(k8sClient))
		v1.GET("/xrs", getXRs(k8sClient))
		v1.GET("/functions", getFunctions(k8sClient))

		// Scope-based endpoints (cluster vs namespace)
		v1.GET("/cluster-resources", getClusterResources(k8sClient))
		v1.GET("/namespace-resources", getNamespaceResources(k8sClient))
	}

	// Serve frontend static files
	router.Static("/_next", "/app/public/_next")
	router.StaticFile("/favicon.ico", "/app/public/favicon.ico")

	// Serve index.html for all other routes (SPA routing)
	router.NoRoute(func(c *gin.Context) {
		c.File("/app/public/index.html")
	})

	return router
}

// corsMiddleware configures CORS for the API
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// healthCheck returns the health status of the API
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "crossplane-spy",
	})
}
