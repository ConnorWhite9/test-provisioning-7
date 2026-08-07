package controller

import (
	"context"

	"go.uber.org/zap"

	"github.com/container-orchestration/system/internal/scheduler"
	"github.com/container-orchestration/system/internal/store"
)

// Manager runs the control loops that reconcile desired state with actual state.
type Manager struct {
	store     store.Store
	scheduler *scheduler.Scheduler
	logger    *zap.Logger
}

// New creates a Manager with the given dependencies.
func New(s store.Store, sched *scheduler.Scheduler, logger *zap.Logger) *Manager {
	return &Manager{store: s, scheduler: sched, logger: logger}
}

// Run starts all control loops and blocks until ctx is cancelled.
func (m *Manager) Run(ctx context.Context) {
	// Replica controller and node health controller will be implemented here.
	<-ctx.Done()
}
