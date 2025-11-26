package repository

import (
	"context"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"
)

func (repo *repository) QuestionDeleteById(ctx context.Context, params entity.QuestionDeleteByIdParams) (entity.QuestionDeleteByIdResult, error) {
	tx := repo.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return entity.QuestionDeleteByIdResult{}, tx.Error
	}
	defer tx.Rollback()

	question := model.Question{QuestionId: params.QuestionId}
	result := tx.Delete(&question)
	if result.Error != nil {
		return entity.QuestionDeleteByIdResult{}, result.Error
	}
	if result.RowsAffected == 0 {
		return entity.QuestionDeleteByIdResult{FoundQuestion: false}, nil
	}

	err := tx.Commit().Error
	if err != nil {
		return entity.QuestionDeleteByIdResult{}, err
	}
	return entity.QuestionDeleteByIdResult{FoundQuestion: true}, nil
}
