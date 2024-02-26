package service

import (
	"quiz/dto"
	"quiz/models"
	"quiz/repository"
	"time"
)

type ServiceQuiz interface {
	GetAllQuiz() ([]*models.Quiz, error)
	CreatedQuiz(input dto.CreatedQuizDto) (*models.Quiz, error)
	UpdatedQuiz(ID int, input dto.UpdatedQuizDto) (*models.Quiz, error)
	DeletedQuiz(ID int) (*models.Quiz, error)
	IsQuizActive() ([]*models.Quiz, error)
}

type quiz_service struct {
	repository_quiz repository.RepositoryQuiz
}

func NewQuizService(repository_quiz repository.RepositoryQuiz) *quiz_service {
	return &quiz_service{repository_quiz}
}

func (s *quiz_service) IsQuizActive() ([]*models.Quiz, error) {
	quizzes, err := s.repository_quiz.GetAllQuiz() // Ambil semua quiz
	if err != nil {
		return nil, err
	}

	activeQuizzes := []*models.Quiz{}
	now := time.Now()

	for _, quiz := range quizzes {
		if quiz.FinishedAt.After(now) {
			activeQuizzes = append(activeQuizzes, quiz)
		}
	}

	return activeQuizzes, nil
}

func (s *quiz_service) GetAllQuiz() ([]*models.Quiz, error) {
	findAllQuiz, err := s.repository_quiz.GetAllQuiz()
	if err != nil {
		return findAllQuiz, err
	}
	return findAllQuiz, nil
}

func (s *quiz_service) DeletedQuiz(ID int) (*models.Quiz, error) {
	findQuiz, err := s.repository_quiz.FindQuizById(ID)
	if err != nil {
		return findQuiz, err
	}

	deletedQuiz, err := s.repository_quiz.DeletedQuizById(findQuiz)
	if err != nil {
		return deletedQuiz, err
	}
	return deletedQuiz, nil
}

func (s *quiz_service) CreatedQuiz(input dto.CreatedQuizDto) (*models.Quiz, error) {
	quiz := &models.Quiz{}

	quiz.Title = input.Title
	quiz.Description = input.Description
	quiz.StartedAt = input.StartedAt
	quiz.FinishedAt = input.FinishedAt

	quizing, err := s.repository_quiz.CreatedQuiz(quiz)
	if err != nil {
		return quizing, err
	}
	return quizing, nil
}

func (s *quiz_service) UpdatedQuiz(ID int, input dto.UpdatedQuizDto) (*models.Quiz, error) {
	// quiz := &models.Quiz{}

	FindQuiz, err := s.repository_quiz.FindQuizById(ID)
	if err != nil {
		return FindQuiz, err
	}

	FindQuiz.Title = input.Title
	FindQuiz.Description = input.Description
	FindQuiz.StartedAt = input.StartedAt
	FindQuiz.FinishedAt = input.FinishedAt

	quizing, err := s.repository_quiz.UpdatedQuiz(FindQuiz)
	if err != nil {
		return quizing, err
	}
	return quizing, nil
}
