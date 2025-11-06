"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * ProviderConfigs page
 */
export default function ProviderConfigsPage() {
  return (
    <ResourceListPage
      title="ProviderConfigs"
      description="Provider configurations for authentication and settings"
      fetchResources={api.getProviderConfigs}
    />
  );
}
