"use client";

import { ResourceListPage } from "@/components/resources/resource-list-page";
import { api } from "@/lib/api";

/**
 * Compositions page
 */
export default function CompositionsPage() {
  return (
    <ResourceListPage
      title="Compositions"
      description="Infrastructure templates that define how to compose managed resources"
      fetchResources={api.getCompositions}
    />
  );
}
