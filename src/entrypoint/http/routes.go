package http

import (
	"github.com/anaclaraddias/feedmaker-students/src/entrypoint/handler"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	Handle(c *gin.Context)
}

const (
	HealthCheck = "healthCheck"
)

type HealthRoutes struct {
	*gin.Engine
	healthHandlers map[string]HandlerInterface
}

func NewHealthRoutes(
	app *gin.Engine,
) *HealthRoutes {
	return &HealthRoutes{
		app,
		createMapOfHealthHandlers(),
	}
}

func (healthRoutes *HealthRoutes) Register() {
	healthRoutes.GET(
		"/health",
		healthRoutes.healthHandlers[HealthCheck].Handle,
	)
}

func createMapOfHealthHandlers() map[string]HandlerInterface {
	return map[string]HandlerInterface{
		HealthCheck: handler.NewHealthHandler(),
	}
}
