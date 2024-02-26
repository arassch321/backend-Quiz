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

type userHandler struct {
	userService service.ServiceUser
	authService auth.Service
}

func NewUserHandler(userService service.ServiceUser, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input dto.RegisterUserDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Periksa ketersediaan email sebelum mendaftarkan pengguna
	isEmailAvailable, err := h.userService.IsEmaillAvailabilty(input.Email)
	if err != nil {
		response := helper.FailedResponse1(http.StatusConflict, err.Error(), input)
		c.JSON(http.StatusConflict, response)
		return
	}

	// Jika email tidak tersedia, kirim respons kesalahan
	if !isEmailAvailable {
		response := helper.FailedResponse1(http.StatusConflict, err.Error(), input)
		c.JSON(http.StatusConflict, response)
		return
	}

	// Register user jika email tersedia
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Format dan kirim respons berhasil jika registrasi berhasil
	response := helper.SuccessfulResponse1(newUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input dto.LoginDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), loggedinUser)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Role)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), loggedinUser)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.SuccessfulResponse1(token)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeletedUser(c *gin.Context) {
	param := c.Param("id")
	params, _ := strconv.Atoi(param)

	_, err := h.userService.DeleteUser(params)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// formatter := user.FormatterUser(newDel, "nil")
	response := helper.APIresponse(http.StatusOK, "Account has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
