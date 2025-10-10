package handler

import (
	"net/http"

	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"github.com/anaclaraddias/feedmaker-students/src/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateFeedbackHandler struct {
	db *gorm.DB
}

func NewCreateFeedbackHandler(db *gorm.DB) HandlerInterface {
	return &CreateFeedbackHandler{db: db}
}

type CreateFeedbackRequest struct {
	Score     int    `json:"score" binding:"required,min=0"`
	Body      string `json:"body" binding:"required"`
	StudentID uint   `json:"student_id" binding:"required"`
	TeacherID uint   `json:"teacher_id" binding:"required"`
}

func (handler *CreateFeedbackHandler) Handle(context *gin.Context) {
	var req CreateFeedbackRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback := &entity.Feedback{
		Score:     req.Score,
		Body:      req.Body,
		StudentID: req.StudentID,
		TeacherID: req.TeacherID,
	}

	if err := service.NewCreateFeedbackService(handler.db).Execute(feedback); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, feedback)
}
