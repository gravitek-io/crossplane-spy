"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";
import {
  Package,
  Settings,
  FileCode2,
  Layers,
  Boxes,
  Home,
  Globe,
  FolderTree,
} from "lucide-react";

/**
 * Navigation item type
 */
interface NavItem {
  title: string;
  href: string;
  icon: React.ComponentType<{ className?: string }>;
  description?: string;
}

/**
 * Navigation sections
 */
const navSections = {
  overview: {
    title: "Overview",
    items: [
      {
        title: "Dashboard",
        href: "/",
        icon: Home,
        description: "Overview of all resources",
      },
    ],
  },
  clusterScope: {
    title: "Cluster Scope",
    items: [
      {
        title: "Providers",
        href: "/providers",
        icon: Package,
        description: "Package installations",
      },
      {
        title: "XRDs",
        href: "/xrds",
        icon: FileCode2,
        description: "Composite resource definitions",
      },
      {
        title: "Compositions",
        href: "/compositions",
        icon: Layers,
        description: "Infrastructure templates",
      },
      {
        title: "Functions",
        href: "/functions",
        icon: Boxes,
        description: "Composition functions",
      },
      {
        title: "ProviderConfigs",
        href: "/providerconfigs",
        icon: Settings,
        description: "Provider configurations",
      },
    ],
  },
  namespaceScope: {
    title: "Namespace Scope",
    items: [
      {
        title: "Composite Resources",
        href: "/xrs",
        icon: FolderTree,
        description: "XR instances",
      },
    ],
  },
};

/**
 * Sidebar navigation component
 */
export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="fixed left-0 top-0 z-40 h-screen w-64 border-r bg-background">
      <div className="flex h-full flex-col">
        {/* Header */}
        <div className="flex h-16 items-center border-b px-6">
          <Link href="/" className="flex items-center gap-2 font-semibold">
            <Globe className="h-6 w-6 text-primary" />
            <span className="text-lg">Crossplane Spy</span>
          </Link>
        </div>

        {/* Navigation */}
        <nav className="flex-1 overflow-y-auto p-4">
          {/* Overview Section */}
          <div className="mb-6">
            <h3 className="mb-2 px-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              {navSections.overview.title}
            </h3>
            <ul className="space-y-1">
              {navSections.overview.items.map((item) => (
                <NavLink
                  key={item.href}
                  item={item}
                  isActive={pathname === item.href}
                />
              ))}
            </ul>
          </div>

          {/* Cluster Scope Section */}
          <div className="mb-6">
            <h3 className="mb-2 px-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              {navSections.clusterScope.title}
            </h3>
            <ul className="space-y-1">
              {navSections.clusterScope.items.map((item) => (
                <NavLink
                  key={item.href}
                  item={item}
                  isActive={pathname === item.href}
                />
              ))}
            </ul>
          </div>

          {/* Namespace Scope Section */}
          <div className="mb-6">
            <h3 className="mb-2 px-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              {navSections.namespaceScope.title}
            </h3>
            <ul className="space-y-1">
              {navSections.namespaceScope.items.map((item) => (
                <NavLink
                  key={item.href}
                  item={item}
                  isActive={pathname === item.href}
                />
              ))}
            </ul>
          </div>
        </nav>

        {/* Footer */}
        <div className="border-t p-4">
          <p className="text-xs text-muted-foreground">
            Educational tool for Crossplane v2
          </p>
        </div>
      </div>
    </aside>
  );
}

/**
 * Navigation link component
 */
function NavLink({ item, isActive }: { item: NavItem; isActive: boolean }) {
  const Icon = item.icon;

  return (
    <li>
      <Link
        href={item.href}
        className={cn(
          "flex items-center gap-3 rounded-md px-3 py-2 text-sm transition-colors",
          isActive
            ? "bg-primary text-primary-foreground"
            : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
        )}
      >
        <Icon className="h-4 w-4" />
        <span>{item.title}</span>
      </Link>
    </li>
  );
}
