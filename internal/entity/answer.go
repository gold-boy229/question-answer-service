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

type AnswerGetByIdParams struct {
	AnswerId int
}

type AnswerGetByIdResult struct {
	Answer      Answer
	FoundAnswer bool
}

type AnswerDeleteByIdParams struct {
	AnswerId int
}

type AnswerDeleteByIdResult struct {
	FoundAnswer bool
}

type Answer struct {
	AnswerId   int
	QuestionId int
	UserId     string
	Text       string
	CreatedAt  string
}
