# Self-Service CI/CD Demo

A deliberately small Go service used as a companion application for the
[self-service-cicd-platform](https://github.com/marco13-moo/self-service-cicd-platform).
The platform can validate this repository, register it in the developer catalog,
and create a preview environment through Argo Workflows.

## Run locally

```bash
go run .
```

The service listens on port `8080`:

```bash
curl http://localhost:8080/
curl http://localhost:8080/healthz
```

Build the container locally:

```bash
docker build -t self-service-cicd-demo:local .
```

## Register with the platform

Start the companion platform first by following its local quickstart. In the
platform repository's `control-plane` directory, set the endpoint and your
local development bearer token in the shell. Do not put credentials in this
repository.

Then run from the platform repository's `control-plane` directory:

```bash
export PLATFORM_ENDPOINT=http://localhost:8080
export PLATFORM_TOKEN=your-local-development-token

cd /path/to/self-service-cicd-platform/control-plane
go run ./cmd/platformctl \
  -endpoint "$PLATFORM_ENDPOINT" \
  -token "$PLATFORM_TOKEN" \
  -file /path/to/self-service-cicd-demo/platform/service.yaml \
  apply
```

Inspect the service:

```bash
go run ./cmd/platformctl \
  -endpoint "$PLATFORM_ENDPOINT" \
  -token "$PLATFORM_TOKEN" \
  catalog

go run ./cmd/platformctl \
  -endpoint "$PLATFORM_ENDPOINT" \
  -token "$PLATFORM_TOKEN" \
  diagnose self-service-cicd-demo
```

## Create a preview environment

After registration, submit a disposable environment through the platform API:

```bash
curl -X POST "$PLATFORM_ENDPOINT/api/v1/environments" \
  -H "Authorization: Bearer $PLATFORM_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "demo-preview",
    "service": "self-service-cicd-demo",
    "ttl": "1h"
  }'
```

Watch the resulting Argo workflows:

```bash
kubectl get workflows -n argo -w
```

The platform owns intent and preview records; Argo owns workflow execution and
lifecycle. This demo declaration has no external egress requirements.
