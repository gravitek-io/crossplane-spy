"use client";

import { useEffect, useState } from "react";
import { Header } from "@/components/nav/header";
import { ResourceTable } from "@/components/resources/resource-table";

/**
 * Generic resource list page component
 */
interface ResourceListPageProps {
  title: string;
  description: string;
  fetchResources: () => Promise<any>;
  columns?: any[];
}

export function ResourceListPage({
  title,
  description,
  fetchResources,
  columns,
}: ResourceListPageProps) {
  const [resources, setResources] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await fetchResources();
        setResources(data.items || []);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to fetch resources");
      } finally {
        setLoading(false);
      }
    };

    fetch();
  }, [fetchResources]);

  return (
    <div className="flex flex-col h-full">
      <Header title={title} description={description} />

      <div className="flex-1 p-6">
        {loading && (
          <div className="flex items-center justify-center h-64">
            <p className="text-muted-foreground">Loading {title.toLowerCase()}...</p>
          </div>
        )}

        {error && (
          <div className="rounded-lg border border-destructive bg-destructive/10 p-4">
            <p className="text-sm text-destructive">{error}</p>
            <p className="text-xs text-muted-foreground mt-2">
              Make sure the backend is running and accessible.
            </p>
          </div>
        )}

        {!loading && !error && (
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <p className="text-sm text-muted-foreground">
                {resources.length} {resources.length === 1 ? "resource" : "resources"} found
              </p>
            </div>
            <ResourceTable resources={resources} columns={columns} />
          </div>
        )}
      </div>
    </div>
  );
}
