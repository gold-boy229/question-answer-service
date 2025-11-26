package handlers

import (
	"context"
	"net/http"
	"question-answer-service/internal/entity"

	"github.com/go-playground/validator/v10"
)

type answerProvider interface {
	AnswerCreate(context.Context, entity.AnswerCreateParams) (entity.AnswerCreateResult, error)
}

type answerHandler struct {
	repo      answerProvider
	validator *validator.Validate
}

func NewAnswerHandler(repo answerProvider, validator *validator.Validate) *answerHandler {
	return &answerHandler{repo: repo, validator: validator}
}

func (h *answerHandler) GetAnswerById(http.ResponseWriter, *http.Request) {}

func (h *answerHandler) DeleteAnswerById(http.ResponseWriter, *http.Request) {}
