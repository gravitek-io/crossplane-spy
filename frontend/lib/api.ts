/**
 * API client for Crossplane Spy backend
 */

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

/**
 * Generic fetch wrapper with error handling
 */
async function fetchAPI<T>(endpoint: string): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`);

  if (!response.ok) {
    throw new Error(`API error: ${response.statusText}`);
  }

  return response.json();
}

/**
 * API client methods
 */
export const api = {
  // Generic resource endpoints
  getResources: () => fetchAPI("/resources"),
  getResourcesByKind: (kind: string) => fetchAPI(`/resources/${kind}`),
  getResource: (kind: string, namespace: string, name: string) =>
    fetchAPI(`/resources/${kind}/${namespace}/${name}`),

  // Specific resource type endpoints
  getProviders: () => fetchAPI("/providers"),
  getProviderConfigs: () => fetchAPI("/providerconfigs"),
  getXRDs: () => fetchAPI("/xrds"),
  getCompositions: () => fetchAPI("/compositions"),
  getXRs: () => fetchAPI("/xrs"),
  getFunctions: () => fetchAPI("/functions"),

  // Scope-based endpoints
  getClusterResources: () => fetchAPI("/cluster-resources"),
  getNamespaceResources: () => fetchAPI("/namespace-resources"),
};
