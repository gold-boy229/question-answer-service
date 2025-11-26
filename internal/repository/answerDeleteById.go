package repository

import (
	"context"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"
)

func (repo *repository) AnswerDeleteById(ctx context.Context, params entity.AnswerDeleteByIdParams) (entity.AnswerDeleteByIdResult, error) {
	answer := model.Answer{AnswerId: params.AnswerId}
	result := repo.DB.WithContext(ctx).Delete(&answer)
	if result.Error != nil {
		return entity.AnswerDeleteByIdResult{}, result.Error
	}
	return entity.AnswerDeleteByIdResult{FoundAnswer: result.RowsAffected > 0}, nil
}
