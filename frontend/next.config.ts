import type { NextConfig } from "next";
import path from "path";

const nextConfig: NextConfig = {
  /* config options here */
  output: "standalone",

  // Explicitly set the workspace root to avoid warnings about multiple lockfiles
  outputFileTracingRoot: path.join(__dirname, ".."),
};

export default nextConfig;
