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
	Questions   []Question
}

type Question struct {
	ID             int
	QuizID         int
	Question       string
	TrueAnswer     string
	Options        string // Opsi jawaban
	StudentAnswers []StudentAnswer
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
