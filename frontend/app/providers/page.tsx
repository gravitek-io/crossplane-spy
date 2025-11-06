"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * Providers page
 */
export default function ProvidersPage() {
  return (
    <ResourceListPage
      title="Providers"
      description="Package installations for extending Crossplane with new capabilities"
      fetchResources={api.getProviders}
    />
  );
}
