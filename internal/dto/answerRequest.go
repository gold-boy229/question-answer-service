package dto

type AnswerCreate_Request struct {
	QuestionId *int   `validate:"required,min=1"`
	UserId     string `json:"user_id" validate:"required"`
	Text       string `json:"text" validate:"required"`
}
