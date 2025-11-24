package model

import "time"

type Question struct {
	Id        int
	Text      string
	CreatedAt time.Time
}
