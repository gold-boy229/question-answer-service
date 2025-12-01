package mocks

import (
	"context"
	"errors"
	"question-answer-service/internal/entity"

	"github.com/stretchr/testify/mock"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// Mock implementation of the questionProvider interface
type QuestionProvider struct {
	mock.Mock
}

func (m *QuestionProvider) QuestionGetAll(ctx context.Context) ([]entity.Question, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionProvider) QuestionCreate(ctx context.Context, params entity.QuestionCreateParams) (entity.QuestionCreateResult, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(entity.QuestionCreateResult), args.Error(1)
}

func (m *QuestionProvider) QuestionGetWithAnswersById(ctx context.Context, params entity.QuestionGetWithAnswersByIdParams) (entity.QuestionGetWithAnswersByIdResult, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(entity.QuestionGetWithAnswersByIdResult), args.Error(1)
}

func (m *QuestionProvider) QuestionDeleteById(ctx context.Context, params entity.QuestionDeleteByIdParams) (entity.QuestionDeleteByIdResult, error) {
	return entity.QuestionDeleteByIdResult{}, ErrNotImplemented
}
