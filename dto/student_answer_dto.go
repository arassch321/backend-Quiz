package dto

type AnswerDto struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}
