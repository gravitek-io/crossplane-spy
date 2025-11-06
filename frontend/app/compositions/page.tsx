"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { ScopeBadge } from "@/components/ui/status-badge";
import { formatDistanceToNow } from "@/lib/date-utils";
import { getResourcePackage } from "@/lib/resource-utils";
import { api } from "@/lib/api";

/**
 * Compositions page
 * Note: Compositions are templates/blueprints and don't have status conditions
 */
export default function CompositionsPage() {
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
      title="Compositions"
      description="Infrastructure templates that define how to compose managed resources"
      fetchResources={api.getCompositions}
      columns={columns}
    />
  );
}
