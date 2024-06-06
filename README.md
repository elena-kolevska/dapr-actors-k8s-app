# README

export GOOS=linux && export DAPR_REGISTRY=diagrid/dapr && export DAPR_TAG=edge && TARGET_OS_LOCAL=linux  && TARGET_OS=linux && export TARGET_ARCH=arm64 && make build && make docker-build


Create kind cluster
```
kind create cluster
```

Pull images
```
docker pull public.ecr.aws/diagrid/dapr:edge
docker pull public.ecr.aws/diagrid/placement:edge-linux-arm64
docker pull public.ecr.aws/diagrid/sentry:edge
docker pull public.ecr.aws/diagrid/operator:edge
docker pull public.ecr.aws/diagrid/injector:edge
```

Load images into kind cluster
```
kind load docker-image actor-server1:latest

kind load docker-image public.ecr.aws/diagrid/dapr:edge
kind load docker-image public.ecr.aws/diagrid/placement:edge
kind load docker-image public.ecr.aws/diagrid/sentry:edge
kind load docker-image public.ecr.aws/diagrid/operator:edge
kind load docker-image public.ecr.aws/diagrid/injector:edge
```

Install Dapr with Helm
```
helm upgrade --install \
    dapr --namespace=dapr-system --wait --timeout 5m0s \
    --set-string global.registry=public.ecr.aws/diagrid \
    --set global.ha.enabled=false \
    --set-string global.tag=edge \
    --set global.logAsJson=true \
    --set global.daprControlPlaneOs=linux \
    --set global.daprControlPlaneArch=arm64 \
    --set dapr_placement.logLevel=debug \
    --set dapr_sidecar_injector.sidecarImagePullPolicy=IfNotPresent \
    --set global.imagePullPolicy=IfNotPresent --set global.imagePullSecrets= \
    --set global.mtls.enabled=true \
    --set dapr_placement.cluster.forceInMemoryLog=true \
    --set dapr_operator.logLevel=debug,dapr_operator.watchInterval=20s --debug \
    /Users/elenakolevska/Diagrid/Code/dapr/charts/dapr/
```

Install Redis
```
helm install redis bitnami/redis
```
Get credentials from Redis (instructions in the output from the install command above), and set them in the state store definitions in the ns1.yaml and ns2.yaml files.