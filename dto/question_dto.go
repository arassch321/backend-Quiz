package dto

type CreatedQuestionDto struct {
	Question   string            `json:"question" binding:"required"`
	TrueAnswer string            `json:"true_answer" binding:"required"`
	Options    map[string]string `json:"options"`
}

type UpdatedQuestionDto struct {
	Question   string            `json:"question" binding:"required"`
	TrueAnswer string            `json:"true_answer" binding:"required"`
	Options    map[string]string `json:"options"`
}
