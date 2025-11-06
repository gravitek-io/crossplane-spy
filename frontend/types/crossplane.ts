/**
 * Common Kubernetes metadata
 */
export interface K8sMetadata {
  name: string;
  namespace?: string;
  labels?: Record<string, string>;
  annotations?: Record<string, string>;
  creationTimestamp?: string;
  uid?: string;
}

/**
 * Kubernetes condition status
 */
export interface K8sCondition {
  type: string;
  status: "True" | "False" | "Unknown";
  lastTransitionTime?: string;
  reason?: string;
  message?: string;
}

/**
 * Base interface for Crossplane resources
 */
export interface CrossplaneResource {
  metadata: K8sMetadata;
  kind: string;
  apiVersion: string;
  spec?: unknown;
  status?: {
    conditions?: K8sCondition[];
    [key: string]: unknown;
  };
}

/**
 * Resource scope types
 */
export type ResourceScope = "cluster" | "namespace";

/**
 * Filter options for resources
 */
export interface ResourceFilters {
  namespace?: string;
  scope?: ResourceScope;
  status?: string;
  search?: string;
}

/**
 * Crossplane resource kinds
 */
export type CrossplaneResourceKind =
  | "Provider"
  | "ProviderConfig"
  | "CompositeResourceDefinition"
  | "Composition"
  | "CompositeResource"
  | "Function";
