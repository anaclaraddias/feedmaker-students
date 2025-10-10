package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) HandlerInterface {
	return &HealthHandler{db: db}
}

func (handler *HealthHandler) Handle(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	sqlDB, err := handler.db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "unhealthy", "error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "unhealthy", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
}
