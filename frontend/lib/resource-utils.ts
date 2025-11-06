/**
 * Utility functions for extracting package/group information from Crossplane resources
 */

/**
 * Extracts the package or group identifier from a resource
 * Different resource types store this information differently:
 * - Providers & Functions: spec.package (full package path)
 * - XRDs: spec.group (API group)
 * - Compositions: spec.compositeTypeRef.apiVersion (contains group/version)
 * - ProviderConfigs & XRs: apiVersion (contains group/version)
 */
export function getResourcePackage(resource: any): string {
  if (!resource) return "";

  // Providers and Functions have spec.package
  if (resource.spec?.package) {
    // Extract organization/package from full path
    // Example: xpkg.crossplane.io/crossplane-contrib/provider-aws-ec2:v2.0.0
    // Returns: crossplane-contrib/provider-aws-ec2
    const pkg = resource.spec.package;
    const match = pkg.match(/^[^/]+\/(.+?)(?::|$)/);
    return match ? match[1] : pkg;
  }

  // XRDs have spec.group
  if (resource.spec?.group) {
    return resource.spec.group;
  }

  // Compositions have spec.compositeTypeRef.apiVersion
  if (resource.spec?.compositeTypeRef?.apiVersion) {
    // Extract group from apiVersion (group/version)
    const apiVersion = resource.spec.compositeTypeRef.apiVersion;
    const group = apiVersion.split("/")[0];
    return group;
  }

  // ProviderConfigs and XRs use apiVersion
  if (resource.apiVersion) {
    // Extract group from apiVersion (group/version)
    const group = resource.apiVersion.split("/")[0];
    // Filter out core groups
    if (group && !group.includes("pkg.crossplane.io") && !group.includes("apiextensions.crossplane.io")) {
      return group;
    }
  }

  return "";
}

/**
 * Gets a short display name for the package
 * For full package paths, returns just the package name
 */
export function getShortPackageName(resource: any): string {
  const fullPackage = getResourcePackage(resource);

  // If it's a full package path (contains /), get the last part
  if (fullPackage.includes("/")) {
    const parts = fullPackage.split("/");
    return parts[parts.length - 1].replace(/:.*$/, ""); // Remove version tag
  }

  return fullPackage;
}
