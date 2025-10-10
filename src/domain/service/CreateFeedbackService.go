package service

import (
	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

type CreateFeedbackService struct {
	db *gorm.DB
}

func NewCreateFeedbackService(db *gorm.DB) *CreateFeedbackService {
	return &CreateFeedbackService{
		db: db,
	}
}

func (service *CreateFeedbackService) Execute(feedback *entity.Feedback) error {
	if err := service.db.Create(feedback).Error; err != nil {
		return err
	}

	return nil
}
