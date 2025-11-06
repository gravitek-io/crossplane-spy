import { StatusBadge, ScopeBadge } from "@/components/ui/status-badge";
import { formatDistanceToNow } from "@/lib/date-utils";

/**
 * Generic resource table component
 */
interface ResourceTableProps {
  resources: any[];
  columns?: {
    key: string;
    label: string;
    render?: (value: any, resource: any) => React.ReactNode;
  }[];
}

export function ResourceTable({ resources, columns }: ResourceTableProps) {
  const defaultColumns = [
    {
      key: "name",
      label: "Name",
      render: (_: any, resource: any) => (
        <div>
          <div className="font-medium">{resource.metadata.name}</div>
          {resource.metadata.namespace && (
            <div className="text-xs text-muted-foreground">
              ns: {resource.metadata.namespace}
            </div>
          )}
        </div>
      ),
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
      key: "status",
      label: "Status",
      render: (_: any, resource: any) => (
        <StatusBadge ready={resource.status?.ready || false} />
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

  const displayColumns = columns || defaultColumns;

  if (resources.length === 0) {
    return (
      <div className="rounded-lg border border-dashed p-8 text-center">
        <p className="text-muted-foreground">No resources found</p>
      </div>
    );
  }

  return (
    <div className="rounded-lg border">
      <div className="overflow-x-auto">
        <table className="w-full">
          <thead className="border-b bg-muted/50">
            <tr>
              {displayColumns.map((column) => (
                <th
                  key={column.key}
                  className="px-4 py-3 text-left text-xs font-medium uppercase tracking-wide text-muted-foreground"
                >
                  {column.label}
                </th>
              ))}
            </tr>
          </thead>
          <tbody className="divide-y">
            {resources.map((resource, index) => (
              <tr
                key={resource.metadata.uid || index}
                className="hover:bg-muted/50 transition-colors"
              >
                {displayColumns.map((column) => (
                  <td key={column.key} className="px-4 py-3">
                    {column.render
                      ? column.render(resource[column.key], resource)
                      : resource[column.key]}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
