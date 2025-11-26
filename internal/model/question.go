package model

import "time"

type Question struct {
	QuestionId int       `gorm:"column:id;primaryKey"`
	Text       string    `gorm:"column:text"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
