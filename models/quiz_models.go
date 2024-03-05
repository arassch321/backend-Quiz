package models

import "time"

type Quiz struct {
	ID          int
	Title       string
	Description string
	StartedAt   time.Time
	FinishedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Questions   []Question `gorm:"constraint:OnDelete:CASCADE"`
}

type Question struct {
	ID             int
	QuizID         int
	Question       string
	TrueAnswer     string
	Options        string
	StudentAnswers []StudentAnswer `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type StudentAnswer struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	QuizID     int
	QuestionID int
	Answer     string
	Score      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
