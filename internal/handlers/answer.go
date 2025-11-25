package handlers

import "net/http"

type answerHandler struct{}

func NewAnswerHandler() *answerHandler {
	return &answerHandler{}
}

func (ah *answerHandler) AddAnswerToQuestion(http.ResponseWriter, *http.Request) {}

func (ah *answerHandler) GetAnswerById(http.ResponseWriter, *http.Request) {}

func (ah *answerHandler) DeleteAnswerById(http.ResponseWriter, *http.Request) {}
