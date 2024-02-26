package service

import (
	"errors"
	"quiz/dto"
	"quiz/models"
	"quiz/repository"
	"time"
)

type ServiceStudentAnswer interface {
	SubmitStudentAnswers(quizID int, userID int, answers []dto.AnswerDto) error
	GetTotalScore(userID int, quizID int) (int, error)
	GetUserAnswersByUserIDAndQuizID(userID, quizID int) ([]models.StudentAnswer, error)
}

type student_answer_service struct {
	repository_student_answer repository.RepositoryStudentAnswer
	repository_question       repository.RepositoryQuestion
}

func NewStudentAnswerService(repository_student_answer repository.RepositoryStudentAnswer, repository_question repository.RepositoryQuestion) *student_answer_service {
	return &student_answer_service{repository_student_answer, repository_question}
}

func (s *student_answer_service) SubmitStudentAnswers(quizID int, userID int, answers []dto.AnswerDto) error {
	// Menghitung bobot per question
	userAnswers, err := s.repository_student_answer.GetUserAnswersByUserIDAndQuizID(userID, quizID)
	if err != nil {
		return err
	}
	if len(userAnswers) > 0 {
		return errors.New("user has already answered the quiz")
	}

	// Menghitung bobot per question
	bobot := 100 / len(answers)

	for _, ans := range answers {
		// Mendapatkan question dari repository atau dari cache jika perlu
		question, err := s.repository_question.GetQuestionByID(ans.QuestionID)
		if err != nil {
			return err
		}

		// Menghitung skor berdasarkan jawaban
		score := 0
		if ans.Answer == question.TrueAnswer {
			score = bobot
		}

		// Membuat objek StudentAnswer
		studentAnswer := &models.StudentAnswer{
			UserID:     userID,
			QuizID:     quizID,
			QuestionID: ans.QuestionID,
			Answer:     ans.Answer,
			Score:      score,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Menyimpan StudentAnswer ke dalam repository
		if err := s.repository_student_answer.CreateStudentAnswer(studentAnswer); err != nil {
			return err
		}
	}
	return nil
}

func (s *student_answer_service) GetTotalScore(userID int, quizID int) (int, error) {
	// Mendapatkan semua jawaban pengguna untuk kuis tertentu
	userAnswers, err := s.repository_student_answer.GetUserAnswersByUserIDAndQuizID(userID, quizID)
	if err != nil {
		return 0, err
	}

	// Menghitung total skor dari semua jawaban
	totalScore := 0
	for _, answer := range userAnswers {
		totalScore += answer.Score
	}

	return totalScore, nil
}

func (s *student_answer_service) GetUserAnswersByUserIDAndQuizID(userID, quizID int) ([]models.StudentAnswer, error) {
	return s.repository_student_answer.GetUserAnswersByUserIDAndQuizID(userID, quizID)
}
