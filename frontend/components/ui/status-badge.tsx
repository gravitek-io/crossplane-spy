import { cn } from "@/lib/utils";
import { CheckCircle2, AlertCircle } from "lucide-react";

/**
 * Status badge component for displaying resource status
 */
interface StatusBadgeProps {
  ready: boolean;
  className?: string;
}

export function StatusBadge({ ready, className }: StatusBadgeProps) {
  return (
    <div
      className={cn(
        "inline-flex items-center gap-1.5 rounded-full px-2.5 py-0.5 text-xs font-medium",
        ready
          ? "bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400"
          : "bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400",
        className
      )}
    >
      {ready ? (
        <CheckCircle2 className="h-3 w-3" />
      ) : (
        <AlertCircle className="h-3 w-3" />
      )}
      <span>{ready ? "Ready" : "Not Ready"}</span>
    </div>
  );
}

/**
 * Scope badge component
 */
interface ScopeBadgeProps {
  scope: "cluster" | "namespace";
  namespace?: string;
  className?: string;
}

export function ScopeBadge({ scope, namespace, className }: ScopeBadgeProps) {
  return (
    <div
      className={cn(
        "inline-flex items-center gap-1 rounded-full px-2.5 py-0.5 text-xs font-medium",
        scope === "cluster"
          ? "bg-blue-100 text-blue-800 dark:bg-blue-900/30 dark:text-blue-400"
          : "bg-purple-100 text-purple-800 dark:bg-purple-900/30 dark:text-purple-400",
        className
      )}
    >
      <span className="font-semibold">{scope === "cluster" ? "Cluster" : "Namespace"}</span>
      {namespace && <span className="text-muted-foreground">· {namespace}</span>}
    </div>
  );
}

/**
 * Detailed status badge for resources with multiple status indicators
 */
interface DetailedStatusBadgeProps {
  status: {
    installed?: boolean;
    healthy?: boolean;
    established?: boolean;
    ready?: boolean;
  };
  className?: string;
}

export function DetailedStatusBadge({ status, className }: DetailedStatusBadgeProps) {
  const statuses = [];

  // Provider/Function statuses
  if (status.installed !== undefined) {
    statuses.push({
      label: "Installed",
      value: status.installed,
    });
  }
  if (status.healthy !== undefined) {
    statuses.push({
      label: "Healthy",
      value: status.healthy,
    });
  }

  // XRD status
  if (status.established !== undefined) {
    statuses.push({
      label: "Established",
      value: status.established,
    });
  }

  // Generic ready status (fallback)
  if (statuses.length === 0 && status.ready !== undefined) {
    statuses.push({
      label: "Ready",
      value: status.ready,
    });
  }

  return (
    <div className={cn("inline-flex items-center gap-1.5", className)}>
      {statuses.map((s, idx) => (
        <div
          key={idx}
          className={cn(
            "inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-xs font-medium",
            s.value
              ? "bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400"
              : "bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400"
          )}
        >
          {s.value ? (
            <CheckCircle2 className="h-3 w-3" />
          ) : (
            <AlertCircle className="h-3 w-3" />
          )}
          <span>{s.label}</span>
        </div>
      ))}
    </div>
  );
}
