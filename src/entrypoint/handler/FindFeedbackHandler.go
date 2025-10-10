package handler

import (
	"net/http"
	"strconv"

	"github.com/anaclaraddias/feedmaker-students/src/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FindFeedbackHandler struct {
	db *gorm.DB
}

func NewFindFeedbackHandler(db *gorm.DB) HandlerInterface {
	return &FindFeedbackHandler{db: db}
}

func (handler *FindFeedbackHandler) Handle(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	queryId := context.Param("id")
	id, err := strconv.ParseUint(queryId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid feedback id"})
		return
	}

	feedback, err := service.NewFindFeedbackService(handler.db).Execute(uint(id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, feedback)
}
