package repository

import (
	"context"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"

	"gorm.io/gorm"
)

func (repo *repository) AnswerGetById(ctx context.Context, params entity.AnswerGetByIdParams) (entity.AnswerGetByIdResult, error) {
	var answer = model.Answer{AnswerId: params.AnswerId}
	err := repo.DB.WithContext(ctx).First(&answer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.AnswerGetByIdResult{FoundAnswer: false}, nil
		}
		return entity.AnswerGetByIdResult{}, err
	}
	return entity.AnswerGetByIdResult{
		FoundAnswer: true,
		Answer:      convertModelToEntity_OneAnswer(answer),
	}, nil
}
