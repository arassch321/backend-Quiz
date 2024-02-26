package dto

import "time"

type CreatedQuizDto struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartedAt   time.Time `json:"startedAt" binding:"required"`
	FinishedAt  time.Time `json:"finishedAt" binding:"required"`
}

type UpdatedQuizDto struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartedAt   time.Time `json:"startedAt" binding:"required"`
	FinishedAt  time.Time `json:"finishedAt" binding:"required"`
}
