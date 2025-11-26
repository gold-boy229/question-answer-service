package handlers

import (
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
)

func (h *questionHandler) DeleteQuestionById(w http.ResponseWriter, req *http.Request) {
	reqDTO, err := bindInputParams_QuestionDeleteById(req)
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

	params := convertDTOToEntity_QuestionDeleteById(reqDTO)
	result, err := h.repo.QuestionDeleteById(req.Context(), params)
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

	w.WriteHeader(http.StatusOK)
}

func bindInputParams_QuestionDeleteById(req *http.Request) (dto.QuestionDeleteById_Request, error) {
	var (
		reqDTO dto.QuestionDeleteById_Request
		err    error
	)

	reqDTO.QuestionId, err = getQuestionIdFromPath(req)
	if err != nil {
		return dto.QuestionDeleteById_Request{}, err
	}
	return reqDTO, nil
}

func convertDTOToEntity_QuestionDeleteById(reqDTO dto.QuestionDeleteById_Request) entity.QuestionDeleteByIdParams {
	return entity.QuestionDeleteByIdParams{
		QuestionId: *reqDTO.QuestionId,
	}
}
