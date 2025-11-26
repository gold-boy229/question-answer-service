package dto

type AnswerCreate_Response struct {
	Answer Answer_Response `json:"answer"`
}

type Answer_Response struct {
	AnswerId   int    `json:"answer_id"`
	QuestionId int    `json:"question_id"`
	UserId     string `json:"user_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}
