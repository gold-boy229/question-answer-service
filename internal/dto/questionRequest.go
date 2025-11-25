package dto

type QuestionCreate_Request struct {
	Text string `json:"text" validate:"required"`
}
