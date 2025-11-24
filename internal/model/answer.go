package model

import "time"

type Answer struct {
	Id         int
	QuestionId int
	UserId     string
	Text       string
	CreatedAt  time.Time
}
