"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * Composite Resources (XRs) page
 */
export default function XRsPage() {
  return (
    <ResourceListPage
      title="Composite Resources (XRs)"
      description="Instances of composite resources created from XRDs"
      fetchResources={api.getXRs}
    />
  );
}
