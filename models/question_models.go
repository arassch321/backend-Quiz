package models

import "time"

type Questions struct {
	ID         int
	QuizID     int
	Question   string
	TrueAnswer string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
