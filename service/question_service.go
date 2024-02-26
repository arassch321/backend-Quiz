package service

import (
	"encoding/json"
	"errors"
	"quiz/dto"
	"quiz/models"
	"quiz/repository"
)

type ServiceQuestion interface {
	AddQuestionsToQuiz(quizID int, questions []dto.CreatedQuestionDto) (*models.Quiz, error)
	GetQuestionsByQuizID(quizID int) ([]*models.Question, error)
	GetQuestionByID(questionID int) (*models.Question, error)
}

type question_service struct {
	repository_question repository.RepositoryQuestion
	repository_quiz     repository.RepositoryQuiz
}

func NewQuestionService(repository_question repository.RepositoryQuestion, repository_quiz repository.RepositoryQuiz) *question_service {
	return &question_service{repository_question, repository_quiz}
}

func (s *question_service) GetQuestionsByQuizID(quizID int) ([]*models.Question, error) {
	questions, err := s.repository_question.GetQuestionsByQuizID(quizID)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("no questions found for the quiz")
	}
	return questions, nil
}

func (s *question_service) AddQuestionsToQuiz(quizID int, questions []dto.CreatedQuestionDto) (*models.Quiz, error) {
	// Get the quiz by ID
	quiz, err := s.repository_quiz.GetQuizByID(quizID)
	if err != nil {
		return nil, err
	}

	// Add questions to the quiz
	for _, q := range questions {
		// Convert options map to JSON
		optionsJSON, err := json.Marshal(q.Options)
		if err != nil {
			return nil, err
		}

		question := &models.Question{
			QuizID:     quiz.ID,
			Question:   q.Question,
			TrueAnswer: q.TrueAnswer,
			Options:    string(optionsJSON),
		}
		err = s.repository_question.CreatedQuestion(question)
		if err != nil {
			return nil, err
		}
		quiz.Questions = append(quiz.Questions, *question)
	}

	return quiz, nil
}

func (s *question_service) GetQuestionByID(questionID int) (*models.Question, error) {
	question, err := s.repository_question.GetQuestionByID(questionID)
	if err != nil {
		return nil, err
	}

	return question, nil
}
