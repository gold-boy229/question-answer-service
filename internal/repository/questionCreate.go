package repository

import (
	"context"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"
)

func (repo *repository) QuestionCreate(ctx context.Context, params entity.QuestionCreateParams) (entity.QuestionCreateResult, error) {
	question := model.Question{
		Text: params.Text,
	}

	result := repo.DB.WithContext(ctx).Create(&question)
	if result.Error != nil {
		return entity.QuestionCreateResult{}, result.Error
	}

	return entity.QuestionCreateResult{
		Question: convertModelToEntity_OneQuestion(question),
	}, nil
}
