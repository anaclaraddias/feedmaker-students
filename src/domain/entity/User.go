package entity

type UserType string

const userStudent UserType = "STUDENT"
const userTeacher UserType = "TEACHER"
const userCoordinator UserType = "COORDINATOR"

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"not null" json:"name"`
	Type     UserType `gorm:"not null" json:"type"`
	Username string   `gorm:"unique;not null" json:"username"`
	Password string   `gorm:"not null" json:"-"`
}

func (User) TableName() string {
	return "user"
}
