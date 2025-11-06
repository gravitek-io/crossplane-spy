# Crossplane Spy - Frontend

Next.js frontend application for Crossplane Spy dashboard.

## Tech Stack

- **Next.js 15** - React framework with App Router
- **TypeScript** - Type safety
- **TailwindCSS** - Utility-first CSS framework
- **shadcn/ui** - Re-usable component library

## Structure

- `app/` - Next.js App Router pages and layouts
- `components/` - React components
- `lib/` - Utility functions and API client
- `types/` - TypeScript type definitions
- `public/` - Static assets

## Development

### Prerequisites

- Node.js 20 or later
- npm or yarn

### Getting Started

```bash
# Install dependencies
npm install

# Run development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) to see the application.

### Building

```bash
# Build for production
npm run build

# Run production build
npm start
```

## Adding shadcn/ui Components

To add a new shadcn/ui component:

```bash
npx shadcn@latest add <component-name>
```

For example:
```bash
npx shadcn@latest add button
npx shadcn@latest add card
npx shadcn@latest add table
```

## Environment Variables

Create a `.env.local` file:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## API Integration

The API client is located in `lib/api.ts` and provides methods to interact with the backend:

- `api.getProviders()` - Fetch all Providers
- `api.getXRDs()` - Fetch all XRDs
- `api.getCompositions()` - Fetch all Compositions
- And more...
