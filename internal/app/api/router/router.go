package router

import (
	"go-gin-boilerplate/internal/app/api/controller"
	"go-gin-boilerplate/internal/app/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router handles all API routing
type Router struct {
	engine           *gin.Engine
	logger           *logrus.Logger
	healthController *controller.HealthController
	noteController   *controller.NoteController
}

// New creates a new router instance
func New(
	logger *logrus.Logger,
	healthController *controller.HealthController,
	noteController *controller.NoteController,
) *Router {
	return &Router{
		engine:           gin.Default(),
		logger:           logger,
		healthController: healthController,
		noteController:   noteController,
	}
}

// Setup configures all routes and middleware
func (r *Router) Setup() *gin.Engine {
	// Add middleware
	r.engine.Use(gin.Recovery())

	// Custom logger
	r.engine.Use(middleware.Logger(r.logger))

	// Custom error handler
	r.engine.Use(middleware.Recover(r.logger))

	// Setup routes
	health := r.engine.Group("/health")
	{
		health.GET("", r.healthController.Check)
	}

	// API v1 group
	v1 := r.engine.Group("/api/v1")
	{
		// Notes endpoints
		notes := v1.Group("/notes")
		{
			notes.POST("", r.noteController.CreateNote)
			notes.GET("", r.noteController.GetAllNotes)
			notes.GET("/:id", r.noteController.GetNote)
			notes.PUT("/:id", r.noteController.UpdateNote)
			notes.DELETE("/:id", r.noteController.DeleteNote)
		}
	}

	return r.engine
}
