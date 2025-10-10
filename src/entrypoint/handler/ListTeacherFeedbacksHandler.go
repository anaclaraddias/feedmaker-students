package handler

import (
	"net/http"
	"strconv"

	"github.com/anaclaraddias/feedmaker-students/src/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListTeacherFeedbacksHandler struct {
	db *gorm.DB
}

func NewListTeacherFeedbacksHandler(db *gorm.DB) HandlerInterface {
	return &ListTeacherFeedbacksHandler{db: db}
}

func (handler *ListTeacherFeedbacksHandler) Handle(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	queryId := context.Param("id")
	id, err := strconv.ParseUint(queryId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid teacher id"})
		return
	}

	feedbacks, err := service.NewListTeacherFeedbackService(handler.db).Execute(uint(id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, feedbacks)
}
