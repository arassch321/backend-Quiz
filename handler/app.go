package handler

import (
	"quiz/auth"
	"quiz/database"
	"quiz/middleware"
	"quiz/repository"
	"quiz/service"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatal("Eror Db Connection")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods:    []string{"POST, OPTIONS, GET, PUT, DELETE"},
	}))

	// user
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)
	//--//
	user := router.Group("/api/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), userHandler.DeletedUser)

	// quiz
	quizRepository := repository.NewQuizRepository(db)
	quizService := service.NewQuizService(quizRepository)
	quizHandler := NewQuizHandler(quizService, authService)
	//--//
	quiz := router.Group("/api/quiz")
	quiz.POST("/created-quiz", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), quizHandler.CreatedQuiz)
	quiz.PUT("/updated-quiz/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), quizHandler.UpdatedQuiz)
	quiz.GET("/", middleware.AuthMiddleware(authService, userService), quizHandler.GetAllQuiz)
	quiz.GET("/is-active", middleware.AuthMiddleware(authService, userService), quizHandler.IsQuizActive)
	quiz.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), quizHandler.DeletedQuiz)

	//question
	questionRepository := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepository, quizRepository)
	questionHandler := NewQuestionHandler(quizService, questionService)
	//--//
	question := router.Group("/api/question")
	question.POST("/created-question/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), questionHandler.AddQuestionsToQuizHandler)
	question.GET("/questions/:quiz_id", questionHandler.GetAllQuestionsByQuizIDHandler)
	question.GET("/:question_id", questionHandler.GetQuestionByIDHandler)

	//student
	studentAnswerRepository := repository.NewStudentAnswerRepository(db)
	studentAnswerService := service.NewStudentAnswerService(studentAnswerRepository, questionRepository)
	studentAnswerHandler := NewStudentAnswerHandler(studentAnswerService, questionService)

	studentAnswer := router.Group("/api/submit-answer")

	studentAnswer.POST("/quizzes/:quizID/submit-answers", middleware.AuthMiddleware(authService, userService), studentAnswerHandler.SubmitStudentAnswersHandler)

	studentAnswer.GET("/api/user/quiz/:quizID/total-score", middleware.AuthMiddleware(authService, userService), studentAnswerHandler.GetTotalScoreHandler)

	// Port
	router.Run(":8080")
}
