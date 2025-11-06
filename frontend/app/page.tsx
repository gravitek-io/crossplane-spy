"use client";

import { useEffect, useState } from "react";
import { Header } from "@/components/nav/header";
import { api } from "@/lib/api";
import { Package, FileCode2, Layers, Boxes, FolderTree, Settings } from "lucide-react";

/**
 * Dashboard overview page
 */
export default function Home() {
  const [summary, setSummary] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchSummary = async () => {
      try {
        const data = await api.getResources();
        setSummary(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to fetch resources");
      } finally {
        setLoading(false);
      }
    };

    fetchSummary();
  }, []);

  return (
    <div className="flex flex-col h-full">
      <Header
        title="Dashboard"
        description="Overview of all Crossplane resources in your cluster"
      />

      <div className="flex-1 p-6">
        {loading && (
          <div className="flex items-center justify-center h-64">
            <p className="text-muted-foreground">Loading resources...</p>
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

        {!loading && !error && summary && (
          <div className="space-y-6">
            <div>
              <h2 className="text-lg font-semibold mb-4">Cluster Resources</h2>
              <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                <ResourceCard
                  title="Providers"
                  count={summary.providers || 0}
                  icon={Package}
                  href="/providers"
                  description="Package installations"
                />
                <ResourceCard
                  title="XRDs"
                  count={summary.xrds || 0}
                  icon={FileCode2}
                  href="/xrds"
                  description="Composite resource definitions"
                />
                <ResourceCard
                  title="Compositions"
                  count={summary.compositions || 0}
                  icon={Layers}
                  href="/compositions"
                  description="Infrastructure templates"
                />
                <ResourceCard
                  title="Functions"
                  count={summary.functions || 0}
                  icon={Boxes}
                  href="/functions"
                  description="Composition functions"
                />
                <ResourceCard
                  title="ProviderConfigs"
                  count={summary.providerConfigs || 0}
                  icon={Settings}
                  href="/providerconfigs"
                  description="Provider configurations"
                />
              </div>
            </div>

            <div>
              <h2 className="text-lg font-semibold mb-4">Namespace Resources</h2>
              <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                <ResourceCard
                  title="Composite Resources"
                  count={summary.compositeResources || 0}
                  icon={FolderTree}
                  href="/xrs"
                  description="XR instances"
                />
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

/**
 * Resource card component
 */
interface ResourceCardProps {
  title: string;
  count: number;
  icon: React.ComponentType<{ className?: string }>;
  href: string;
  description: string;
}

function ResourceCard({ title, count, icon: Icon, href, description }: ResourceCardProps) {
  return (
    <a
      href={href}
      className="block rounded-lg border bg-card p-6 shadow-sm transition-all hover:shadow-md hover:border-primary"
    >
      <div className="flex items-center justify-between mb-4">
        <Icon className="h-8 w-8 text-primary" />
        <span className="text-3xl font-bold">{count}</span>
      </div>
      <h3 className="text-lg font-semibold mb-1">{title}</h3>
      <p className="text-sm text-muted-foreground">{description}</p>
    </a>
  );
}
