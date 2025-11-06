"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { ScopeBadge } from "@/components/ui/status-badge";
import { formatDistanceToNow } from "@/lib/date-utils";
import { getResourcePackage } from "@/lib/resource-utils";
import { api } from "@/lib/api";

/**
 * ProviderConfigs page
 * Note: ProviderConfigs may not have meaningful status conditions depending on the provider
 */
export default function ProviderConfigsPage() {
  const columns = [
    {
      key: "name",
      label: "Name",
      render: (_: any, resource: any) => (
        <div className="font-medium">{resource.metadata.name}</div>
      ),
    },
    {
      key: "package",
      label: "Package / Group",
      render: (_: any, resource: any) => {
        const pkg = getResourcePackage(resource);
        return pkg ? (
          <span className="text-sm font-mono text-muted-foreground">{pkg}</span>
        ) : (
          <span className="text-sm text-muted-foreground/50">-</span>
        );
      },
    },
    {
      key: "scope",
      label: "Scope",
      render: (_: any, resource: any) => (
        <ScopeBadge
          scope={resource.scope}
          namespace={resource.metadata.namespace}
        />
      ),
    },
    {
      key: "age",
      label: "Age",
      render: (_: any, resource: any) => (
        <span className="text-sm text-muted-foreground">
          {formatDistanceToNow(resource.metadata.creationTimestamp)}
        </span>
      ),
    },
  ];

  return (
    <ResourceListPage
      title="ProviderConfigs"
      description="Provider configurations for authentication and settings"
      fetchResources={api.getProviderConfigs}
      columns={columns}
    />
  );
}
