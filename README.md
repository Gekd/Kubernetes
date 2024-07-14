# K8 ClusterF\*

An e-commerce service deployable to a Kubernetes.

## Commands

### Set up Minikube and Open Dashboard

### Create Minikube Docker Env

```powershell
minikube docker-env
```

### Enter Minikube Docker Env

```powershell
& minikube -p minikube docker-env --shell powershell | Invoke-Expression
```

### See Active Services

```powershell
kubectl get services
```

### Build Docker Image

```powershell
docker build -t <image-name> .
```

### Run Pod

```powershell
kubectl run <pod-name> --image=<image-name> --image-pull-policy=Never --restart=Never
```

### Expose Pod Port

```powershell
kubectl expose pod <pod-name> --type=LoadBalancer --port=3000 --target-port=3000
```

### Edit Running Service Configuration (i.e Change port) (opens Text editor)

```powershell
kubectl edit service product-service-container
```

### Create Minikube Tunnel (gives Loadbalancer Services an External Ip)

see [[#See Active Services]] to get the external ip

```powershell
minikube tunnel
```
