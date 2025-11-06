# Crossplane Spy Helm Chart

This Helm chart deploys Crossplane Spy, a dashboard for visualizing Crossplane v2 resources.

## Prerequisites

- Kubernetes 1.20+
- Helm 3.0+
- Crossplane v2 installed in the cluster

## Installation

### Add the repository (when published)

```bash
helm repo add crossplane-spy https://gravitek.github.io/crossplane-spy
helm repo update
```

### Install from source

```bash
# From the project root
helm install crossplane-spy ./helm/crossplane-spy
```

### Install with custom values

```bash
helm install crossplane-spy ./helm/crossplane-spy \
  --set image.repository=your-registry/crossplane-spy \
  --set image.tag=v0.1.0
```

## Configuration

The following table lists the configurable parameters of the Crossplane Spy chart and their default values.

| Parameter | Description | Default |
|-----------|-------------|---------|
| `namespace.create` | Create dedicated namespace | `true` |
| `namespace.name` | Namespace name | `crossplane-spy` |
| `image.repository` | Image repository | `crossplane-spy` |
| `image.tag` | Image tag | `latest` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `replicaCount` | Number of replicas | `1` |
| `serviceAccount.create` | Create service account | `true` |
| `rbac.create` | Create RBAC resources | `true` |
| `service.type` | Service type | `ClusterIP` |
| `service.port` | Service port | `8080` |
| `ingress.enabled` | Enable ingress | `false` |
| `resources.limits.cpu` | CPU limit | `500m` |
| `resources.limits.memory` | Memory limit | `512Mi` |
| `resources.requests.cpu` | CPU request | `100m` |
| `resources.requests.memory` | Memory request | `128Mi` |

## Accessing the Dashboard

### Using kubectl port-forward

```bash
kubectl port-forward -n crossplane-spy svc/crossplane-spy 8080:8080
```

Then access the dashboard at `http://localhost:8080`.

### Using Ingress

Enable ingress and configure your ingress controller:

```yaml
ingress:
  enabled: true
  className: nginx
  hosts:
    - host: crossplane-spy.example.com
      paths:
        - path: /
          pathType: Prefix
  tls:
    - secretName: crossplane-spy-tls
      hosts:
        - crossplane-spy.example.com
```

## RBAC

By default, the chart creates:
- A dedicated namespace (`crossplane-spy`)
- A ServiceAccount
- A ClusterRole with read-only access to Crossplane resources
- A ClusterRoleBinding

The application only requires **read-only** access to:
- Providers, Functions (pkg.crossplane.io)
- XRDs, Compositions (apiextensions.crossplane.io)
- All custom resources (for discovering XR instances)

## Uninstallation

```bash
helm uninstall crossplane-spy
```

To also delete the namespace:

```bash
kubectl delete namespace crossplane-spy
```

## Security

The application runs with:
- Non-root user (UID 1000)
- Read-only root filesystem
- Dropped capabilities
- Security context constraints

## Troubleshooting

### Pods not starting

Check the pod status and logs:

```bash
kubectl get pods -n crossplane-spy
kubectl logs -n crossplane-spy deployment/crossplane-spy
```

### RBAC issues

Verify the ClusterRole and ClusterRoleBinding are created:

```bash
kubectl get clusterrole crossplane-spy
kubectl get clusterrolebinding crossplane-spy
```

### No resources shown

Ensure Crossplane is properly installed:

```bash
kubectl get providers
kubectl get xrds
```

## Support

For issues and questions:
- GitHub Issues: https://github.com/gravitek/crossplane-spy/issues
- Documentation: https://github.com/gravitek/crossplane-spy
