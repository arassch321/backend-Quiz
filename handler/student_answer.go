package handler

import (
	"net/http"
	"quiz/dto"
	"quiz/models"
	"quiz/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type studentAnswerHandler struct {
	studentAnswerService service.ServiceStudentAnswer
	questionService      service.ServiceQuestion
}

func NewStudentAnswerHandler(studentAnswerService service.ServiceStudentAnswer, questionService service.ServiceQuestion) *studentAnswerHandler {
	return &studentAnswerHandler{studentAnswerService, questionService}
}

func (h *studentAnswerHandler) SubmitStudentAnswersHandler(c *gin.Context) {
	// Mengambil user dari context
	currentUser := c.MustGet("currentUser").(*models.User)
	userID := currentUser.ID

	// Mendapatkan ID Quiz dari URL
	quizID, err := strconv.Atoi(c.Param("quizID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	// Binding JSON request body ke struct
	var answers []struct {
		QuestionID int    `json:"questionID"`
		Answer     string `json:"answer"`
	}
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Konversi answers ke tipe []dto.AnswerDto
	var answerDtos []dto.AnswerDto
	for _, ans := range answers {
		answerDto := dto.AnswerDto{
			QuestionID: ans.QuestionID,
			Answer:     ans.Answer,
		}
		answerDtos = append(answerDtos, answerDto)
	}

	// Memanggil fungsi SubmitStudentAnswers dengan tipe yang sesuai
	err = h.studentAnswerService.SubmitStudentAnswers(quizID, userID, answerDtos)

	// 	// Iterasi setiap jawaban dan kirim ke service
	// 	for _, ans := range answers {
	// err := h.studentAnswerService.SubmitStudentAnswers(quizID, userID, answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User answers submitted successfully"})
}

func (h *studentAnswerHandler) GetTotalScoreHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*models.User)
	userID := currentUser.ID

	quizID, err := strconv.Atoi(c.Param("quizID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	// Mendapatkan jawaban pengguna
	userAnswers, err := h.studentAnswerService.GetUserAnswersByUserIDAndQuizID(userID, quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total skor dan jawaban yang benar
	totalScore := 0
	correctAnswers := 0
	var correctAnswersDetails []gin.H
	for _, answer := range userAnswers {
		totalScore += answer.Score
		correctAnswers++
		question, err := h.questionService.GetQuestionByID(answer.QuestionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		correctAnswersDetails = append(correctAnswersDetails, gin.H{
			"questionID":    answer.QuestionID,
			"yourAnswer":    answer.Answer,
			"correctAnswer": question.TrueAnswer,
		})
	}

	// Membuat respons
	response := gin.H{
		"totalScore":            totalScore,
		"correctAnswers":        correctAnswers,
		"correctAnswersDetails": correctAnswersDetails,
	}

	c.JSON(http.StatusOK, response)
}
