package handler

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Handle(c *gin.Context)
}

type HealthHandler struct {
}

func NewHealthHandler() HandlerInterface {
	return &HealthHandler{}
}

func (healthHandler *HealthHandler) Handle(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeader(200)

	context.JSON(200, map[string]interface{}{"status": "healthy"})
}
