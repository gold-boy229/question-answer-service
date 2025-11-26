package repository

import (
	"context"
	"question-answer-service/internal/consts"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/model"

	"gorm.io/gorm"
)

func (repo *repository) QuestionGetWithAnswersById(ctx context.Context, params entity.QuestionGetWithAnswersByIdParams) (entity.QuestionGetWithAnswersByIdResult, error) {
	tx := repo.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return entity.QuestionGetWithAnswersByIdResult{}, tx.Error
	}
	defer tx.Rollback()

	question := model.Question{QuestionId: params.QuestionId}
	err := tx.First(&question).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.QuestionGetWithAnswersByIdResult{FoundQuestion: false}, nil
		}
		return entity.QuestionGetWithAnswersByIdResult{}, err
	}

	answers, err := getQuestionAnswers(tx, params.QuestionId)
	if err != nil {
		return entity.QuestionGetWithAnswersByIdResult{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return entity.QuestionGetWithAnswersByIdResult{}, err
	}

	return entity.QuestionGetWithAnswersByIdResult{
		FoundQuestion: true,
		Question:      convertModelToEntity_OneQuestion(question),
		ShortAnswers:  convertModelToEntity_ManyShortAnswers(answers),
	}, nil
}

func getQuestionAnswers(tx *gorm.DB, questionId int) ([]model.Answer, error) {
	var answers []model.Answer
	err := tx.Where("question_id = ?", questionId).Find(&answers).Error
	return answers, err
}

func convertModelToEntity_ManyShortAnswers(answers []model.Answer) []entity.AnswerShort {
	result := make([]entity.AnswerShort, 0, len(answers))
	for _, ans := range answers {
		result = append(result, convertModelToEntity_OneShortAnswer(ans))
	}
	return result
}

func convertModelToEntity_OneShortAnswer(ans model.Answer) entity.AnswerShort {
	return entity.AnswerShort{
		AnswerBaseFields: entity.AnswerBaseFields{
			AnswerId:  ans.AnswerId,
			UserId:    ans.UserId,
			Text:      ans.Text,
			CreatedAt: ans.CreatedAt.Format(consts.FORMAT_LAYOUT_DATE_TIME),
		},
	}
}
