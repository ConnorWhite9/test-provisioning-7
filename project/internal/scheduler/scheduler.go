package scheduler

import (
	"context"
	"errors"

	"github.com/container-orchestration/system/internal/store"
)

// ErrNoEligibleNode is returned when no node can satisfy the resource request.
var ErrNoEligibleNode = errors.New("no eligible node available for scheduling")

// Scheduler selects a node for a workload instance.
type Scheduler struct {
	store store.Store
}

// New creates a Scheduler backed by the given store.
func New(s store.Store) *Scheduler {
	return &Scheduler{store: s}
}

// Schedule selects the best node for the given workload and returns its ID.
// Selection strategy: filter nodes that are online and have sufficient capacity,
// then pick the node with the most available resources (spread load).
func (s *Scheduler) Schedule(ctx context.Context, workload store.Workload) (string, error) {
	nodes, err := s.store.ListNodes(ctx)
	if err != nil {
		return "", err
	}

	var best *store.Node
	var bestAvailCPU int64

	for i := range nodes {
		n := &nodes[i]
		if n.Status != store.NodeStatusOnline {
			continue
		}
		availCPU := n.CPUCapacity - n.CPUAllocated
		availMem := n.MemoryCapacity - n.MemoryAllocated
		if availCPU < workload.CPURequest || availMem < workload.MemoryRequest {
			continue
		}
		if best == nil || availCPU > bestAvailCPU {
			best = n
			bestAvailCPU = availCPU
		}
	}

	if best == nil {
		return "", ErrNoEligibleNode
	}
	return best.ID, nil
}
