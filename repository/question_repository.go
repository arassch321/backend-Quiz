package repository

import (
	"errors"
	"quiz/models"

	"gorm.io/gorm"
)

type RepositoryQuestion interface {
	GetAllQuestion() ([]*models.Question, error)
	CreatedQuestion(question *models.Question) error
	UpdatedQuestion(question *models.Question) (*models.Question, error)
	DeletedQuestionById(question *models.Question) (*models.Question, error)
	GetQuestionsByQuizID(quizID int) ([]*models.Question, error)
	GetQuestionByID(questionID int) (*models.Question, error)
}

type question_repository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *question_repository {
	return &question_repository{db}
}

func (r *question_repository) CreatedQuestion(question *models.Question) error {
	err := r.db.Create(&question).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *question_repository) GetAllQuestion() ([]*models.Question, error) {
	var question []*models.Question
	err := r.db.Find(&question).Error
	if err != nil {
		return question, err
	}
	return question, nil
}

func (r *question_repository) DeletedQuestionById(question *models.Question) (*models.Question, error) {
	err := r.db.Delete(&question).Error
	if err != nil {
		return question, err
	}
	return question, nil
}

func (r *question_repository) GetQuestionByID(questionID int) (*models.Question, error) {
	var question models.Question
	if err := r.db.First(&question, questionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &question, nil
}

func (r *question_repository) GetQuestionsByQuizID(quizID int) ([]*models.Question, error) {
	var questions []*models.Question
	if err := r.db.Where("quiz_id = ?", quizID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *question_repository) UpdatedQuestion(question *models.Question) (*models.Question, error) {
	err := r.db.Save(&question).Error
	if err != nil {
		return question, err
	}

	return question, nil
}
