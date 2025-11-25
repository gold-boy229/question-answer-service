package handlers

import (
	"context"
	"net/http"
	"question-answer-service/internal/entity"
)

type questionProvider interface {
	GetAllQuestions(context.Context) ([]entity.Question, error)
}

type questionHandler struct {
	repo questionProvider
}

func NewQuestionHandler(repo questionProvider) *questionHandler {
	return &questionHandler{repo: repo}
}

func (qh *questionHandler) CreateQuestion(http.ResponseWriter, *http.Request) {}

func (qh *questionHandler) GetQuestionWithAnswersById(http.ResponseWriter, *http.Request) {}

func (qh *questionHandler) DeleteQuestionById(http.ResponseWriter, *http.Request) {}
