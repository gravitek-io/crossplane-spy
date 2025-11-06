package models

import "time"

// ResourceScope defines whether a resource is cluster or namespace scoped
type ResourceScope string

const (
	ScopeCluster   ResourceScope = "cluster"
	ScopeNamespace ResourceScope = "namespace"
)

// Metadata represents common Kubernetes metadata
type Metadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace,omitempty"`
	UID               string            `json:"uid"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp time.Time         `json:"creationTimestamp"`
}

// Condition represents a Kubernetes status condition
type Condition struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"` // True, False, Unknown
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Message            string    `json:"message,omitempty"`
}

// ResourceStatus represents common status fields
type ResourceStatus struct {
	Conditions []Condition `json:"conditions,omitempty"`
	Ready      bool        `json:"ready"`
}

// BaseResource represents common fields for all Crossplane resources
type BaseResource struct {
	Kind       string        `json:"kind"`
	APIVersion string        `json:"apiVersion"`
	Metadata   Metadata      `json:"metadata"`
	Scope      ResourceScope `json:"scope"`
}

// ResourceReference represents a reference to another resource
type ResourceReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}
