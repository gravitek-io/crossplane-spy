"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * Functions page
 */
export default function FunctionsPage() {
  return (
    <ResourceListPage
      title="Functions"
      description="Composition functions for advanced resource transformation"
      fetchResources={api.getFunctions}
    />
  );
}
