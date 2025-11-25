package handlers

import "net/http"

type questionHandler struct{}

func NewQuestionHandler() *questionHandler {
	return &questionHandler{}
}

func (qh *questionHandler) GetAllQuestions(http.ResponseWriter, *http.Request) {}

func (qh *questionHandler) CreateQuestion(http.ResponseWriter, *http.Request) {}

func (qh *questionHandler) GetQuestionWithAnswersById(http.ResponseWriter, *http.Request) {}

func (qh *questionHandler) DeleteQuestionById(http.ResponseWriter, *http.Request) {}
