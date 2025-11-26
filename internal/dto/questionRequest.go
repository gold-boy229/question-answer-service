package dto

type QuestionCreate_Request struct {
	Text string `json:"text" validate:"required"`
}

type QuestionGetById_Request struct {
	QuestionId *int `validate:"required,min=1"`
}
