# K8S Ingress API

## Requirements

- Kubernetes >= 1.19
- An ingress controller installed (defaults to nginx Ingress Controller)
- Optional: Cert Manager

## Install & Run

### Create a ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: ingress-api-config
  namespace: default
data:
  config.yml: |
    ingress:
      service:
        name: example-svc
        port:
          number: 80
```

For more information about possible configuration keys, [read below](#configuration).

### Create Service Account, role and Role Binding

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: ingress-api
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: manage-ingress
rules:
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get", "list", "create", "update", "patch", "delete"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: manage-ingress-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: manage-ingress
subjects:
- apiGroup: ""
  kind: ServiceAccount
  name: ingress-api
```

### Deploy the API

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ingress-api
  labels:
    name: ingress-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: ingress-api
  template:
    metadata:
      labels:
        app: ingress-api
    spec:
      serviceAccountName: ingress-api
      containers:
        - name: ingress-api
          image: kumojin/ingress-api:v1.0.0
          imagePullPolicy: Always
          ports:
            - containerPort: 3000
          readinessProbe:
            httpGet:
              path: /ping
              port: 3000
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            httpGet:
              path: /ping
              port: 3000
            initialDelaySeconds: 5
            periodSeconds: 5
          volumeMounts:
            - mountPath: /etc/k8s-ingress-api
              name: config
              readOnly: true
      volumes:
        - name: config
          configMap:
            name: ingress-api-config
---
apiVersion: v1
kind: Service
metadata:
  name: ingress-api
spec:
  ports:
    - name: http
      targetPort: 3000
      port: 80
  selector:
    app: ingress-api
```

## Usage

| Verb | Route                            | Query Params     | Note                                                                                                  |
|------|----------------------------------|------------------|-------------------------------------------------------------------------------------------------------|
| GET  | `/ping`                          |                  | Always answers `pong` with 200 a status                                                                |
| POST | `/ingress`                       | `dryRun`, `host` | Creates a new Ingress linked to the `host`. Returns the created Ingress Spec as returned by Kubernetes  |
| GET  | `/cname/:cname/matches/:matches` |                  | Checks if `:cname` matches the given `:matches`, useful before creating an ingress with a TLS certificate.          |

## Configuration

The application is looking for a `config.yml` file located in the `/etc/k8s-ingress-api` directory or in the current application directory.

| Key                           | Required | Default   | Notes                                                                                                              |
|-------------------------------|----------|-----------|--------------------------------------------------------------------------------------------------------------------|
| `namespace`                   | Yes      | `default` | Namespace used for the ingress.                                                                                      |
| `ingress.class`               | Yes      | `nginx`   | Value of `kubernetes.io/ingress.class`.                                                                             |
| `ingress.customMeta`          | No       | `[]`      | List of custom metadata to add to the Ingress spec like `"nginx.ingress.kubernetes.io/force-ssl-redirect": "true"` |
| `ingress.service.name`        | Yes      |           | Service Name used as backend for the ingress.                                                                       |
| `ingress.service.port.number` | Yes*     |           | Service port number used as backend for the ingress.                                                               |
| `ingress.service.port.name`   | Yes*     |           | Service port name as backend for the ingress.                                                                       |

> 📓 **Note**: If `ingress.service.port.number` is not set, the application falls back to `ingress.service.port.name` and if none of them are set, the application does not start.
