package model

import "time"

type Question struct {
	Id        int `gorm:"primaryKey"`
	Text      string
	CreatedAt time.Time
}
