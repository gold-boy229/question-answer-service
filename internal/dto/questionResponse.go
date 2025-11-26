package dto

type Question_Response struct {
	Id        int    `json:"question_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type QuestionGetById_Response struct {
	Question     Question_Response      `json:"question"`
	ShortAnswers []AnswerShort_Response `json:"answers"`
}
