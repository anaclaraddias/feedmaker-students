package handler

import (
	"net/http"
	"strconv"

	"github.com/anaclaraddias/feedmaker-students/src/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListStudentFeedbacksHandler struct {
	db *gorm.DB
}

func NewListStudentFeedbacksHandler(db *gorm.DB) HandlerInterface {
	return &ListStudentFeedbacksHandler{db: db}
}

func (handler *ListStudentFeedbacksHandler) Handle(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	queryId := context.Param("id")
	id, err := strconv.ParseUint(queryId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	feedbacks, err := service.NewListStudentFeedbackService(handler.db).Execute(uint(id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, feedbacks)
}
