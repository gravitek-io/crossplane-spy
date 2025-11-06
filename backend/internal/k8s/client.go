package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client wraps Kubernetes clients for interacting with the cluster
type Client struct {
	// Clientset is the standard Kubernetes client
	Clientset kubernetes.Interface
	// DynamicClient is used for working with custom resources
	DynamicClient dynamic.Interface
	// Config is the Kubernetes REST config
	Config *rest.Config
}

// NewClient creates a new Kubernetes client
// It attempts to create an in-cluster config first, falling back to kubeconfig
func NewClient() (*Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// Not running in cluster, try kubeconfig
		config, err = getKubeconfigConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get Kubernetes config: %w", err)
		}
	}

	// Create standard Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes clientset: %w", err)
	}

	// Create dynamic client for custom resources
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	return &Client{
		Clientset:     clientset,
		DynamicClient: dynamicClient,
		Config:        config,
	}, nil
}

// getKubeconfigConfig attempts to load kubeconfig from standard locations
func getKubeconfigConfig() (*rest.Config, error) {
	// Check KUBECONFIG environment variable
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		// Default to ~/.kube/config
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		kubeconfigPath = filepath.Join(homeDir, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build config from kubeconfig: %w", err)
	}

	return config, nil
}
