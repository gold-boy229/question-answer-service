package model

import "time"

type Answer struct {
	AnswerId   int       `gorm:"column:id;primarykey"`
	QuestionId int       `gorm:"column:question_id"`
	UserId     string    `gorm:"column:user_id"`
	Text       string    `gorm:"column:text"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
