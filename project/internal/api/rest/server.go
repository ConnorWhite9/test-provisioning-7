package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/container-orchestration/system/internal/store"
)

// Server is the REST API server for the control plane.
type Server struct {
	store  store.Store
	logger *zap.Logger
	router *gin.Engine
}

// New creates a REST Server.
func New(s store.Store, logger *zap.Logger) *Server {
	router := gin.New()
	router.Use(gin.Recovery())

	srv := &Server{store: s, logger: logger, router: router}
	srv.registerRoutes()
	return srv
}

// registerRoutes wires URL paths to handler functions.
func (s *Server) registerRoutes() {
	v1 := s.router.Group("/api/v1")

	v1.GET("/healthz", s.handleHealthz)

	// Workload routes
	v1.POST("/workloads", s.handleCreateWorkload)
	v1.GET("/workloads", s.handleListWorkloads)
	v1.GET("/workloads/:name", s.handleGetWorkload)
	v1.PUT("/workloads/:name", s.handleUpdateWorkload)
	v1.DELETE("/workloads/:name", s.handleDeleteWorkload)

	// Node routes
	v1.GET("/nodes", s.handleListNodes)
	v1.GET("/nodes/:id", s.handleGetNode)

	// Log proxy
	v1.GET("/workloads/:name/logs", s.handleGetLogs)
}

// Handler stubs — full implementations will be added in feature tasks.

func (s *Server) handleHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) handleCreateWorkload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleListWorkloads(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleGetWorkload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleUpdateWorkload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleDeleteWorkload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleListNodes(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleGetNode(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) handleGetLogs(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ServeHTTP implements http.Handler so Server can be used with net/http.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Handler returns the underlying gin router as an http.Handler.
func (s *Server) Handler() http.Handler {
	return s.router
}
