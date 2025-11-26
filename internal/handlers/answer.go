package handlers

import (
	"context"
	"question-answer-service/internal/entity"

	"github.com/go-playground/validator/v10"
)

type answerProvider interface {
	AnswerCreate(context.Context, entity.AnswerCreateParams) (entity.AnswerCreateResult, error)
	AnswerGetById(context.Context, entity.AnswerGetByIdParams) (entity.AnswerGetByIdResult, error)
	AnswerDeleteById(context.Context, entity.AnswerDeleteByIdParams) (entity.AnswerDeleteByIdResult, error)
}

type answerHandler struct {
	repo      answerProvider
	validator *validator.Validate
}

func NewAnswerHandler(repo answerProvider, validator *validator.Validate) *answerHandler {
	return &answerHandler{repo: repo, validator: validator}
}
