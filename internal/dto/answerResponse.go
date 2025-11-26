package dto

type AnswerCreate_Response struct {
	Answer Answer_Response `json:"answer"`
}

type AnswerGetById_Response struct {
	Answer Answer_Response `json:"answer"`
}

type Answer_Response struct {
	AnswerBaseFields_Response
	QuestionId int `json:"question_id"`
}

type AnswerShort_Response struct {
	AnswerBaseFields_Response
}

type AnswerBaseFields_Response struct {
	AnswerId  int    `json:"answer_id"`
	UserId    string `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
