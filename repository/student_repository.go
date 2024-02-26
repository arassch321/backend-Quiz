package repository

import (
	"quiz/models"

	"gorm.io/gorm"
)

type RepositoryStudentAnswer interface {
	CreateStudentAnswer(studentAnswer *models.StudentAnswer) error
	GetUserAnswersByUserIDAndQuizID(userID int, quizID int) ([]models.StudentAnswer, error)
}

type student_answer_repository struct {
	db *gorm.DB
}

func NewStudentAnswerRepository(db *gorm.DB) *student_answer_repository {
	return &student_answer_repository{db}
}

func (r *student_answer_repository) CreateStudentAnswer(studentAnswer *models.StudentAnswer) error {
	if err := r.db.Create(studentAnswer).Error; err != nil {
		return err
	}
	return nil
}

func (r *student_answer_repository) GetUserAnswersByUserIDAndQuizID(userID int, quizID int) ([]models.StudentAnswer, error) {
	var userAnswers []models.StudentAnswer
	if err := r.db.Where("user_id = ? AND quiz_id = ?", userID, quizID).Find(&userAnswers).Error; err != nil {
		return nil, err
	}
	return userAnswers, nil
}
