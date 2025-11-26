package dto

type QuestionGetAll_Response struct {
	Questions []Question_Response `json:"questions"`
}

type QuestionGetById_Response struct {
	Question     Question_Response      `json:"question"`
	ShortAnswers []AnswerShort_Response `json:"answers"`
}

type Question_Response struct {
	Id        int    `json:"question_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
