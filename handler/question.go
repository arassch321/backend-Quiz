package handler

import (
	"net/http"
	"quiz/dto"
	"quiz/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService service.ServiceQuestion
	quizService     service.ServiceQuiz
}

func NewQuestionHandler(quizService service.ServiceQuiz, questionService service.ServiceQuestion) *questionHandler {
	return &questionHandler{questionService, quizService}
}

func (h *questionHandler) GetAllQuestionsByQuizIDHandler(c *gin.Context) {
	quizIDStr := c.Param("quiz_id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	questions, err := h.questionService.GetQuestionsByQuizID(quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": questions})
}

func (h *questionHandler) GetQuestionByIDHandler(c *gin.Context) {
	questionIDStr := c.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	question, err := h.questionService.GetQuestionByID(questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if question == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"question": question})
}

func (h *questionHandler) AddQuestionsToQuizHandler(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	var questions []dto.CreatedQuestionDto
	if err := c.BindJSON(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk menambahkan pertanyaan ke dalam Quiz
	quiz, err := h.questionService.AddQuestionsToQuiz(quizID, questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"quiz": quiz})
}

func (h *questionHandler) DeleteQuestionByIDHandler(c *gin.Context) {
	questionIDStr := c.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	question, err := h.questionService.DeleteQuestionByID(questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if question == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"question": question})
}
