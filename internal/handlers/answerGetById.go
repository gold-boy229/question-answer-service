package handlers

import (
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
	"question-answer-service/internal/routes"
	"strconv"
)

func (h *answerHandler) GetAnswerById(w http.ResponseWriter, req *http.Request) {
	reqDTO, err := bindInputParams_AnswerGetById(req)
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

	params := convertDTOToEntity_AnswerGetById(reqDTO)
	result, err := h.repo.AnswerGetById(req.Context(), params)
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

	ResponseWithJSON(w, http.StatusOK,
		dto.AnswerGetById_Response{
			Answer: convertEntityToDTO_OneAnswer(result.Answer),
		},
	)
}

func bindInputParams_AnswerGetById(req *http.Request) (dto.AnswerGetById_Request, error) {
	var (
		reqDTO dto.AnswerGetById_Request
		err    error
	)

	reqDTO.AnswerId, err = getAnswerIdFromPath(req)
	if err != nil {
		return dto.AnswerGetById_Request{}, err
	}
	return reqDTO, nil
}

func getAnswerIdFromPath(req *http.Request) (*int, error) {
	answerIdStr := req.PathValue(routes.PathAnswerId)
	val, err := strconv.Atoi(answerIdStr)
	if err != nil {
		return nil, ErrInvalidPathParameter
	}
	return &val, nil
}

func convertDTOToEntity_AnswerGetById(reqDTO dto.AnswerGetById_Request) entity.AnswerGetByIdParams {
	return entity.AnswerGetByIdParams{
		AnswerId: *reqDTO.AnswerId,
	}
}
