package entity

type Feedback struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Score     int    `gorm:"not null" json:"score"`
	Body      string `gorm:"type:text;not null" json:"body"`
	StudentID uint   `gorm:"not null" json:"student_id"`
	TeacherID uint   `gorm:"not null" json:"teacher_id"`
}

func (Feedback) TableName() string {
	return "feedback"
}
