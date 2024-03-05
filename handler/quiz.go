package handler

import (
	"net/http"
	"quiz/auth"
	"quiz/dto"
	"quiz/helper"
	"quiz/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type quizHandler struct {
	quizService service.ServiceQuiz
	authService auth.Service
}

func NewQuizHandler(quizService service.ServiceQuiz, authService auth.Service) *quizHandler {
	return &quizHandler{quizService, authService}
}

func (h *quizHandler) GetAllQuiz(c *gin.Context) {
	findAllQuiz, err := h.quizService.GetAllQuiz()
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), findAllQuiz)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(findAllQuiz)
	c.JSON(http.StatusOK, response)

}

func (h *quizHandler) DeletedQuiz(c *gin.Context) {
	param := c.Param("id")
	params, _ := strconv.Atoi(param)

	deletedQuiz, err := h.quizService.DeletedQuiz(params)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), deletedQuiz)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(deletedQuiz)
	c.JSON(http.StatusOK, response)
}

func (h *quizHandler) IsQuizActive(c *gin.Context) {
	isActive, err := h.quizService.IsQuizActive()
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), isActive)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(isActive)
	c.JSON(http.StatusOK, response)
}

func (h *quizHandler) CreatedQuiz(c *gin.Context) {
	var input dto.CreatedQuizDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	createQuiz, err := h.quizService.CreatedQuiz(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), createQuiz)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(createQuiz)
	c.JSON(http.StatusOK, response)
}

func (h *quizHandler) UpdatedQuiz(c *gin.Context) {
	param := c.Param("id")
	params, _ := strconv.Atoi(param)

	var input dto.UpdatedQuizDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	createQuiz, err := h.quizService.UpdatedQuiz(params, input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), createQuiz)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(createQuiz)
	c.JSON(http.StatusOK, response)
}

func (h *quizHandler) GetQuizByID(c *gin.Context) {
	param := c.Param("id")
	params, _ := strconv.Atoi(param)

	quiz, err := h.quizService.GetQuizByID(params)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), quiz)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(quiz)
	c.JSON(http.StatusOK, response)
}
