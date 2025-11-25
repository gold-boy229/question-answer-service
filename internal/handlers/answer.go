package handlers

import "net/http"

type answerHandler struct{}

func NewAnswerHandler() *answerHandler {
	return &answerHandler{}
}

func (h *answerHandler) AddAnswerToQuestion(http.ResponseWriter, *http.Request) {}

func (h *answerHandler) GetAnswerById(http.ResponseWriter, *http.Request) {}

func (h *answerHandler) DeleteAnswerById(http.ResponseWriter, *http.Request) {}
