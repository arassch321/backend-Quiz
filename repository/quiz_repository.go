package repository

import (
	"quiz/models"

	"gorm.io/gorm"
)

type RepositoryQuiz interface {
	GetAllQuiz() ([]*models.Quiz, error)
	CreatedQuiz(quiz *models.Quiz) (*models.Quiz, error)
	UpdatedQuiz(quiz *models.Quiz) (*models.Quiz, error)
	FindQuizById(ID int) (*models.Quiz, error)
	DeletedQuizById(quiz *models.Quiz) (*models.Quiz, error)
	GetQuizByID(quizID int) (*models.Quiz, error)
}

type quiz_repository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *quiz_repository {
	return &quiz_repository{db}
}

func (r *quiz_repository) GetQuizByID(quizID int) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := r.db.First(&quiz, quizID).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *quiz_repository) GetAllQuiz() ([]*models.Quiz, error) {
	var quizing []*models.Quiz
	err := r.db.Find(&quizing).Error
	if err != nil {
		return quizing, err
	}
	return quizing, nil
}

func (r *quiz_repository) DeletedQuizById(quiz *models.Quiz) (*models.Quiz, error) {
	err := r.db.Delete(&quiz).Error
	if err != nil {
		return quiz, err
	}

	return quiz, nil
}

func (r *quiz_repository) FindQuizById(ID int) (*models.Quiz, error) {
	var quiz *models.Quiz

	err := r.db.Where("id = ?", ID).Find(&quiz).Error

	if err != nil {
		return quiz, err
	}
	return quiz, nil
}

func (r *quiz_repository) CreatedQuiz(quiz *models.Quiz) (*models.Quiz, error) {
	err := r.db.Create(&quiz).Error

	if err != nil {
		return quiz, err
	}
	return quiz, nil
}

func (r *quiz_repository) UpdatedQuiz(quiz *models.Quiz) (*models.Quiz, error) {
	err := r.db.Save(&quiz).Error
	if err != nil {
		return quiz, err
	}

	return quiz, nil
}
