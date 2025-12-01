package mocks

import (
	"context"
	"question-answer-service/internal/entity"

	"github.com/stretchr/testify/mock"
)

type AnswerProvider struct {
	mock.Mock
}

func (m *AnswerProvider) AnswerCreate(ctx context.Context, params entity.AnswerCreateParams) (entity.AnswerCreateResult, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(entity.AnswerCreateResult), args.Error(1)
}

func (m *AnswerProvider) AnswerGetById(ctx context.Context, params entity.AnswerGetByIdParams) (entity.AnswerGetByIdResult, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(entity.AnswerGetByIdResult), args.Error(1)
}

func (m *AnswerProvider) AnswerDeleteById(ctx context.Context, params entity.AnswerDeleteByIdParams) (entity.AnswerDeleteByIdResult, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(entity.AnswerDeleteByIdResult), args.Error(1)
}
