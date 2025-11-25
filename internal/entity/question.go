package entity

type QuestionCreateParams struct {
	Text string
}

type QuestionCreateResult struct {
	Question Question
}

type Question struct {
	Id        int
	Text      string
	CreatedAt string
}
