package handlers

import (
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
)

func (h *answerHandler) DeleteAnswerById(w http.ResponseWriter, req *http.Request) {
	reqDTO, err := bindInputParams_AnswerDeleteById(req)
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

	params := convertDTOToEntity_AnswerDelete(reqDTO)
	result, err := h.repo.AnswerDeleteById(req.Context(), params)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_INTERNAL_SERVER_ERROR, err.Error()))
		return
	}
	if !result.FoundAnswer {
		ResponseWithJSON(w, http.StatusNotFound,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_NOT_FOUND, "Ответ не найден"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func bindInputParams_AnswerDeleteById(req *http.Request) (dto.AnswerDeleteById_Request, error) {
	var (
		reqDTO dto.AnswerDeleteById_Request
		err    error
	)

	reqDTO.AnswerId, err = getAnswerIdFromPath(req)
	if err != nil {
		return dto.AnswerDeleteById_Request{}, err
	}
	return reqDTO, nil
}

func convertDTOToEntity_AnswerDelete(reqDTO dto.AnswerDeleteById_Request) entity.AnswerDeleteByIdParams {
	return entity.AnswerDeleteByIdParams{
		AnswerId: *reqDTO.AnswerId,
	}
}
