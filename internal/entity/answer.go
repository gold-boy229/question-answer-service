package entity

type AnswerCreateParams struct {
	QuestionId int
	UserId     string
	Text       string
}

type AnswerCreateResult struct {
	Answer        Answer
	FoundQuestion bool
}

type Answer struct {
	AnswerId   int
	QuestionId int
	UserId     string
	Text       string
	CreatedAt  string
}
