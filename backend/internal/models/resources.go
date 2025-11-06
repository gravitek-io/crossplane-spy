package models

// Provider represents a Crossplane Provider resource
type Provider struct {
	BaseResource
	Status ProviderStatus `json:"status"`
	Spec   ProviderSpec   `json:"spec,omitempty"`
}

type ProviderSpec struct {
	Package string `json:"package"`
}

type ProviderStatus struct {
	ResourceStatus
	CurrentRevision string `json:"currentRevision,omitempty"`
	InstalledBundle string `json:"installedBundle,omitempty"`
}

// ProviderConfig represents a provider configuration
type ProviderConfig struct {
	BaseResource
	Status ResourceStatus `json:"status"`
	// Spec varies by provider, kept as raw for flexibility
}

// XRD represents a CompositeResourceDefinition
type XRD struct {
	BaseResource
	Status XRDStatus `json:"status"`
	Spec   XRDSpec   `json:"spec,omitempty"`
}

type XRDSpec struct {
	Group            string   `json:"group"`
	ClaimNames       *Names   `json:"claimNames,omitempty"`
	CompositeNames   Names    `json:"compositeNames"`
	DefaultCompositeDeletePolicy string `json:"defaultCompositeDeletePolicy,omitempty"`
}

type Names struct {
	Kind     string `json:"kind"`
	Plural   string `json:"plural"`
	Singular string `json:"singular,omitempty"`
}

type XRDStatus struct {
	ResourceStatus
	Controllers Controllers `json:"controllers,omitempty"`
}

type Controllers struct {
	CompositeResourceClaimTypeRef TypeReference `json:"compositeResourceClaimTypeRef,omitempty"`
	CompositeResourceTypeRef      TypeReference `json:"compositeResourceTypeRef,omitempty"`
}

type TypeReference struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

// Composition represents a Crossplane Composition
type Composition struct {
	BaseResource
	Status CompositionStatus `json:"status"`
	Spec   CompositionSpec   `json:"spec,omitempty"`
}

type CompositionSpec struct {
	CompositeTypeRef TypeReference `json:"compositeTypeRef"`
	Mode             string        `json:"mode,omitempty"` // Pipeline or Resources
	PipelineCount    int           `json:"pipelineCount,omitempty"`
	ResourcesCount   int           `json:"resourcesCount,omitempty"`
}

type CompositionStatus struct {
	ResourceStatus
}

// Function represents a Crossplane Function
type Function struct {
	BaseResource
	Status FunctionStatus `json:"status"`
	Spec   FunctionSpec   `json:"spec,omitempty"`
}

type FunctionSpec struct {
	Package string `json:"package"`
}

type FunctionStatus struct {
	ResourceStatus
	CurrentRevision string `json:"currentRevision,omitempty"`
}

// CompositeResource represents a composite resource (XR) instance
type CompositeResource struct {
	BaseResource
	Status CompositeResourceStatus `json:"status"`
	Spec   CompositeResourceSpec   `json:"spec,omitempty"`
}

type CompositeResourceSpec struct {
	CompositionRef      *ResourceReference `json:"compositionRef,omitempty"`
	CompositionSelector *map[string]string `json:"compositionSelector,omitempty"`
	ResourceRefs        []ResourceReference `json:"resourceRefs,omitempty"`
}

type CompositeResourceStatus struct {
	ResourceStatus
	CompositionRef *ResourceReference  `json:"compositionRef,omitempty"`
	ResourceRefs   []ResourceReference `json:"resourceRefs,omitempty"`
}

// ResourceList represents a list of resources with metadata
type ResourceList struct {
	Kind  string         `json:"kind"`
	Count int            `json:"count"`
	Scope ResourceScope  `json:"scope"`
	Items []BaseResource `json:"items"`
}

// ResourceSummary provides a summary view of all resources
type ResourceSummary struct {
	Providers       int `json:"providers"`
	ProviderConfigs int `json:"providerConfigs"`
	XRDs            int `json:"xrds"`
	Compositions    int `json:"compositions"`
	Functions       int `json:"functions"`
	CompositeResources int `json:"compositeResources"`
}
