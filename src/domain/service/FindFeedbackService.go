package service

import (
	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"gorm.io/gorm"
)

type FindFeedbackService struct {
	db *gorm.DB
}

func NewFindFeedbackService(db *gorm.DB) *FindFeedbackService {
	return &FindFeedbackService{
		db: db,
	}
}

func (service *FindFeedbackService) Execute(id uint) (*entity.Feedback, error) {
	var feedback *entity.Feedback
	if err := service.db.Where("id = ?", id).Find(&feedback).Error; err != nil {
		return nil, err
	}

	return feedback, nil
}
