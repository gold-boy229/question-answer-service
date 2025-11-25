package app

import "net/http"

type questionHandler interface {
	GetAllQuestions(http.ResponseWriter, *http.Request)
	CreateQuestion(http.ResponseWriter, *http.Request)
	GetQuestionWithAnswersById(http.ResponseWriter, *http.Request)
	DeleteQuestionById(http.ResponseWriter, *http.Request)
}

type answerHandler interface {
	AddAnswerToQuestion(http.ResponseWriter, *http.Request)
	GetAnswerById(http.ResponseWriter, *http.Request)
	DeleteAnswerById(http.ResponseWriter, *http.Request)
}
