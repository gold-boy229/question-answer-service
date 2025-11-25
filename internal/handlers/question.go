package handlers

import (
	"context"
	"net/http"
	"question-answer-service/internal/entity"

	"github.com/go-playground/validator/v10"
)

type questionProvider interface {
	QuestionGetAll(context.Context) ([]entity.Question, error)
	QuestionCreate(context.Context, entity.QuestionCreateParams) (entity.QuestionCreateResult, error)
}

type questionHandler struct {
	repo      questionProvider
	validator *validator.Validate
}

func NewQuestionHandler(repo questionProvider, validator *validator.Validate) *questionHandler {
	return &questionHandler{repo: repo, validator: validator}
}

func (h *questionHandler) GetQuestionWithAnswersById(http.ResponseWriter, *http.Request) {}

func (h *questionHandler) DeleteQuestionById(http.ResponseWriter, *http.Request) {}
