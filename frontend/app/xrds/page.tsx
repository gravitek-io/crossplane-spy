"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * XRDs (Composite Resource Definitions) page
 */
export default function XRDsPage() {
  return (
    <ResourceListPage
      title="Composite Resource Definitions (XRDs)"
      description="Define the schema and API for composite resources"
      fetchResources={api.getXRDs}
    />
  );
}
