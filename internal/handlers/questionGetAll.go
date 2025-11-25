package handlers

import (
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
)

func (h *questionHandler) GetAllQuestions(w http.ResponseWriter, req *http.Request) {
	questions, err := h.repo.QuestionGetAll(req.Context())
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_INTERNAL_SERVER_ERROR, err.Error()))
	}

	ResponseWithJSON(w, http.StatusOK, convertEntityToDTO_ManyQuestions(questions))
}

func convertEntityToDTO_ManyQuestions(questions []entity.Question) []dto.Question_Response {
	result := make([]dto.Question_Response, 0, len(questions))
	for _, q := range questions {
		result = append(result, convertEntityToDTO_OneQuestion(q))
	}
	return result
}

func convertEntityToDTO_OneQuestion(q entity.Question) dto.Question_Response {
	return dto.Question_Response{
		Id:        q.Id,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
	}
}
