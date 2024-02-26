package models

import "time"

type StudentAnswerss struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	QuizID     int
	QuestionID int
	Answer     string
	Score      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
