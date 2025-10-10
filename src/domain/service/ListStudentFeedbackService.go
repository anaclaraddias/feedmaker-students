package service

import (
	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

type ListStudentFeedbackService struct {
	db *gorm.DB
}

func NewListStudentFeedbackService(db *gorm.DB) *ListStudentFeedbackService {
	return &ListStudentFeedbackService{
		db: db,
	}
}

func (service *ListStudentFeedbackService) Execute(studentId uint) ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := service.db.Where("student_id = ?", studentId).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return feedbacks, nil
}
