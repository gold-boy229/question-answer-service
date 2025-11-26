package repository

import (
	"context"
	"question-answer-service/internal/consts"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"

	"gorm.io/gorm"
)

func (repo *repository) AnswerCreate(ctx context.Context, params entity.AnswerCreateParams) (entity.AnswerCreateResult, error) {
	tx := repo.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return entity.AnswerCreateResult{}, tx.Error
	}
	defer tx.Rollback()

	questionExists, err := doesQuestionExists(tx, params.QuestionId)
	if err != nil {
		return entity.AnswerCreateResult{}, err
	}
	if !questionExists {
		return entity.AnswerCreateResult{FoundQuestion: false}, nil
	}

	answer := convertEntityToModel_Answer(params)
	if err := tx.Create(&answer).Error; err != nil {
		return entity.AnswerCreateResult{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entity.AnswerCreateResult{}, err
	}

	return entity.AnswerCreateResult{
		Answer:        convertModelToEntity_OneAnswer(answer),
		FoundQuestion: true,
	}, nil
}

func doesQuestionExists(tx *gorm.DB, questionId int) (bool, error) {
	question := model.Question{QuestionId: questionId}
	if err := tx.First(&question).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func convertEntityToModel_Answer(params entity.AnswerCreateParams) model.Answer {
	return model.Answer{
		QuestionId: params.QuestionId,
		UserId:     params.UserId,
		Text:       params.Text,
	}
}

func convertModelToEntity_OneAnswer(answer model.Answer) entity.Answer {
	return entity.Answer{
		AnswerId:   answer.AnswerId,
		QuestionId: answer.QuestionId,
		UserId:     answer.UserId,
		Text:       answer.Text,
		CreatedAt:  answer.CreatedAt.Format(consts.FORMAT_LAYOUT_DATE_TIME),
	}
}
