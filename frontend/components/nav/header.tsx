/**
 * Header component for the main content area
 */
interface HeaderProps {
  title: string;
  description?: string;
  actions?: React.ReactNode;
}

export function Header({ title, description, actions }: HeaderProps) {
  return (
    <div className="border-b bg-background">
      <div className="flex h-16 items-center justify-between px-6">
        <div>
          <h1 className="text-2xl font-bold tracking-tight">{title}</h1>
          {description && (
            <p className="text-sm text-muted-foreground">{description}</p>
          )}
        </div>
        {actions && <div>{actions}</div>}
      </div>
    </div>
  );
}
