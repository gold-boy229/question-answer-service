package entity

type QuestionCreateParams struct {
	Text string
}

type QuestionCreateResult struct {
	Question Question
}

type QuestionGetWithAnswersByIdParams struct {
	QuestionId int
}

type QuestionGetWithAnswersByIdResult struct {
	Question      Question
	ShortAnswers  []AnswerShort
	FoundQuestion bool
}

type Question struct {
	Id        int
	Text      string
	CreatedAt string
}
