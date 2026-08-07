package store

import "context"

// Store defines the interface for cluster state storage.
// All control plane components interact with cluster state through this interface.
type Store interface {
	// Node operations
	PutNode(ctx context.Context, node Node) error
	GetNode(ctx context.Context, id string) (Node, error)
	ListNodes(ctx context.Context) ([]Node, error)
	DeleteNode(ctx context.Context, id string) error

	// Workload operations
	PutWorkload(ctx context.Context, workload Workload) error
	GetWorkload(ctx context.Context, name string) (Workload, error)
	ListWorkloads(ctx context.Context) ([]Workload, error)
	DeleteWorkload(ctx context.Context, name string) error

	// Instance operations
	PutInstance(ctx context.Context, instance Instance) error
	GetInstance(ctx context.Context, id string) (Instance, error)
	ListInstances(ctx context.Context) ([]Instance, error)
	ListInstancesByWorkload(ctx context.Context, workloadName string) ([]Instance, error)
	DeleteInstance(ctx context.Context, id string) error

	// Watch — notifies caller of changes under the given key prefix.
	Watch(ctx context.Context, prefix string) (<-chan WatchEvent, error)

	// Close releases underlying resources.
	Close() error
}

// NodeStatus represents the health state of a node.
type NodeStatus string

const (
	NodeStatusOnline  NodeStatus = "online"
	NodeStatusOffline NodeStatus = "offline"
)

// Node represents a worker node registered with the control plane.
type Node struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Address        string     `json:"address"`
	Status         NodeStatus `json:"status"`
	CPUCapacity    int64      `json:"cpu_capacity"`    // millicores
	MemoryCapacity int64      `json:"memory_capacity"` // bytes
	CPUAllocated   int64      `json:"cpu_allocated"`   // millicores
	MemoryAllocated int64     `json:"memory_allocated"`
	LastHeartbeat  int64      `json:"last_heartbeat"` // Unix timestamp
	Version        int64      `json:"version"`        // optimistic concurrency
}

// Workload represents a user-defined application definition.
type Workload struct {
	Name           string `json:"name"`
	Image          string `json:"image"`
	DesiredReplicas int32  `json:"desired_replicas"`
	CPURequest     int64  `json:"cpu_request"`    // millicores
	MemoryRequest  int64  `json:"memory_request"` // bytes
	Version        int64  `json:"version"`
}

// InstanceStatus represents the runtime state of a container instance.
type InstanceStatus string

const (
	InstanceStatusPending  InstanceStatus = "pending"
	InstanceStatusRunning  InstanceStatus = "running"
	InstanceStatusStopped  InstanceStatus = "stopped"
	InstanceStatusFailed   InstanceStatus = "failed"
	InstanceStatusUnknown  InstanceStatus = "unknown"
)

// Instance represents a single running container of a workload.
type Instance struct {
	ID           string         `json:"id"`
	WorkloadName string         `json:"workload_name"`
	NodeID       string         `json:"node_id"`
	ContainerID  string         `json:"container_id"`
	Status       InstanceStatus `json:"status"`
	Version      int64          `json:"version"`
}

// WatchEventType indicates whether a key was created, modified, or deleted.
type WatchEventType string

const (
	WatchEventPut    WatchEventType = "put"
	WatchEventDelete WatchEventType = "delete"
)

// WatchEvent carries a single state-change notification from the store.
type WatchEvent struct {
	Type  WatchEventType
	Key   string
	Value []byte
}
