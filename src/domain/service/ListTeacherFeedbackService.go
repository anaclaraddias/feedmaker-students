package service

import (
	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

type ListTeacherFeedbackService struct {
	db *gorm.DB
}

func NewListTeacherFeedbackService(db *gorm.DB) *ListTeacherFeedbackService {
	return &ListTeacherFeedbackService{
		db: db,
	}
}

func (service *ListTeacherFeedbackService) Execute(teacherId uint) ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := service.db.Where("teacher_id = ?", teacherId).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return feedbacks, nil
}
