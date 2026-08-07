# Container Orchestration System

A simplified container orchestration system inspired by Kubernetes, implemented in Go.
Manages containerized workloads across a cluster of worker nodes with automated deployment,
scaling, and basic self-healing.

## Architecture

```
┌───────────────────── Control Plane ─────────────────────┐
│                                                           │
│  ┌──────────┐   ┌───────────┐   ┌────────────────────┐  │
│  │ REST API │   │   gRPC    │   │ Controller Manager │  │
│  │  :8080   │   │ API :9090 │   │  (control loops)   │  │
│  └────┬─────┘   └─────┬─────┘   └─────────┬──────────┘  │
│       │               │                   │              │
│       └───────────────┴───────────────────┘              │
│                           │                              │
│                  ┌────────┴────────┐                     │
│                  │  State Store    │                      │
│                  │  (etcd embed)   │                      │
│                  └────────────────┘                      │
└───────────────────────────────────────────────────────────┘
           ▲ gRPC              ▲ gRPC
           │                  │
    ┌──────┴──────┐    ┌──────┴──────┐
    │ Node Agent  │    │ Node Agent  │
    │  worker-1   │    │  worker-2   │
    │  (Docker)   │    │  (Docker)   │
    └─────────────┘    └─────────────┘
```

## Components

| Binary | Description |
|--------|-------------|
| `bin/control-plane` | Central management: API servers, scheduler, controller manager, etcd |
| `bin/agent` | Worker node daemon: Docker management, heartbeat, status reporting |
| `bin/orchctl` | CLI for operators |

## Prerequisites

- Go 1.22+
- Docker daemon running
- `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc` (for gRPC codegen)

## Quick Start

```bash
# Install protoc plugins (one-time)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate gRPC code, then build all binaries
make all

# Run the control plane (uses defaults from environment or .env.example)
./bin/control-plane

# Run a node agent (on a worker node)
AGENT_NODE_NAME=worker-1 AGENT_ADVERTISE_ADDR=192.168.1.10 ./bin/agent

# Use the CLI
./bin/orchctl get nodes
./bin/orchctl create workload --name nginx-app --image nginx:latest --replicas 3
./bin/orchctl get workloads
./bin/orchctl scale workload nginx-app --replicas 5
```

## Environment Variables

See [`.env.example`](.env.example) for the full list of configuration variables.

## Project Layout

```
.
├── cmd/
│   ├── control-plane/   # Control plane entry point
│   ├── agent/           # Node agent entry point
│   └── orchctl/         # CLI entry point
├── internal/
│   ├── store/           # State store interface and types
│   ├── scheduler/       # Node selection logic
│   ├── controller/      # Replica and node health control loops
│   ├── api/
│   │   └── rest/        # Gin REST API server
│   └── proto/           # Generated gRPC code (from make proto)
└── proto/               # .proto source files
```

## Build Phases

1. **Phase 1 — State Store:** etcd-backed implementation of `store.Store`
2. **Phase 2 — Control Plane API:** Complete REST handlers + gRPC server
3. **Phase 3 — Node Agent:** Docker integration, heartbeat, status reporting
4. **Phase 4 — Controller Manager:** Replica and node health control loops
5. **Phase 5 — CLI:** `orchctl` subcommands
6. **Phase 6 — Networking:** VXLAN overlay, IP allocation, iptables port exposure
