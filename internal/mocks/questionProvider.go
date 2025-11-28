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

func (_m *QuestionProvider) QuestionCreate(ctx context.Context, params entity.QuestionCreateParams) (entity.QuestionCreateResult, error) {
	return entity.QuestionCreateResult{}, ErrNotImplemented
}

func (_m *QuestionProvider) QuestionGetWithAnswersById(ctx context.Context, params entity.QuestionGetWithAnswersByIdParams) (entity.QuestionGetWithAnswersByIdResult, error) {
	return entity.QuestionGetWithAnswersByIdResult{}, ErrNotImplemented
}

func (_m *QuestionProvider) QuestionDeleteById(ctx context.Context, params entity.QuestionDeleteByIdParams) (entity.QuestionDeleteByIdResult, error) {
	return entity.QuestionDeleteByIdResult{}, ErrNotImplemented
}
