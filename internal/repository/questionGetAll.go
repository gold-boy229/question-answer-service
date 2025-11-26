package repository

import (
	"context"
	"question-answer-service/internal/consts"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"
)

func (repo *repository) QuestionGetAll(ctx context.Context) ([]entity.Question, error) {
	questions := make([]model.Question, 0)
	_ = repo.DB.Find(&questions)

	return convertModelToEntity_ManyQuestions(questions), nil
}

func convertModelToEntity_ManyQuestions(questions []model.Question) []entity.Question {
	result := make([]entity.Question, 0, len(questions))
	for _, q := range questions {
		result = append(result, convertModelToEntity_OneQuestion(q))
	}
	return result
}

func convertModelToEntity_OneQuestion(q model.Question) entity.Question {
	return entity.Question{
		Id:        q.QuestionId,
		Text:      q.Text,
		CreatedAt: q.CreatedAt.Format(consts.FORMAT_LAYOUT_DATE_TIME),
	}
}
