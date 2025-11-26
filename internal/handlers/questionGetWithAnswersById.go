package handlers

import (
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
)

func (h *questionHandler) GetQuestionWithAnswersById(w http.ResponseWriter, req *http.Request) {
	reqDTO, err := bindInputParams_QuestionGetById(req)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_BAD_REQUEST, err.Error()),
		)
		return
	}
	if err = h.validator.Struct(reqDTO); err != nil {
		ResponseWithJSON(w, http.StatusBadRequest,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_BAD_REQUEST, err.Error()),
		)
		return
	}

	params := convertDTOToEntity_QuestionGetByIdWithAnswers(reqDTO)
	result, err := h.repo.QuestionGetWithAnswersById(req.Context(), params)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_INTERNAL_SERVER_ERROR, err.Error()))
		return
	}
	if !result.FoundQuestion {
		ResponseWithJSON(w, http.StatusNotFound,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_NOT_FOUND, "Вопрос не найден"))
		return
	}

	ResponseWithJSON(w, http.StatusOK,
		dto.QuestionGetById_Response{
			Question:     convertEntityToDTO_OneQuestion(result.Question),
			ShortAnswers: convertEntityToDTO_ManyShortAnswers(result.ShortAnswers),
		},
	)
}

func bindInputParams_QuestionGetById(req *http.Request) (dto.QuestionGetById_Request, error) {
	var (
		reqDTO dto.QuestionGetById_Request
		err    error
	)

	reqDTO.QuestionId, err = getQuestionIdFromPath(req)
	if err != nil {
		return dto.QuestionGetById_Request{}, err
	}
	return reqDTO, nil
}

func convertDTOToEntity_QuestionGetByIdWithAnswers(reqDTO dto.QuestionGetById_Request) entity.QuestionGetWithAnswersByIdParams {
	return entity.QuestionGetWithAnswersByIdParams{
		QuestionId: *reqDTO.QuestionId,
	}
}

func convertEntityToDTO_ManyShortAnswers(shortAnswers []entity.AnswerShort) []dto.AnswerShort_Response {
	result := make([]dto.AnswerShort_Response, 0, len(shortAnswers))
	for _, shortAns := range shortAnswers {
		result = append(result, convertEntityToDTO_OneShortAnswer(shortAns))
	}
	return result
}

func convertEntityToDTO_OneShortAnswer(shortAns entity.AnswerShort) dto.AnswerShort_Response {
	return dto.AnswerShort_Response{
		AnswerBaseFields_Response: dto.AnswerBaseFields_Response{
			AnswerId:  shortAns.AnswerId,
			UserId:    shortAns.UserId,
			Text:      shortAns.Text,
			CreatedAt: shortAns.CreatedAt,
		},
	}
}
