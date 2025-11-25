package handlers

import (
	"encoding/json"
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
)

func (h *questionHandler) CreateQuestion(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var reqDTO dto.QuestionCreate_Request
	if err := json.NewDecoder(req.Body).Decode(&reqDTO); err != nil {
		ResponseWithJSON(w, http.StatusBadRequest,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_BAD_REQUEST, err.Error()),
		)
		return
	}
	if err := h.validator.Struct(reqDTO); err != nil {
		ResponseWithJSON(w, http.StatusBadRequest,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_BAD_REQUEST, err.Error()),
		)
		return
	}

	params := convertDTOToEntity_QuestionCreate(reqDTO)
	result, err := h.repo.QuestionCreate(req.Context(), params)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError,
			dto.NewErrorResponse(enum.ERROR_RESPONSE_INTERNAL_SERVER_ERROR, err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusCreated, convertEntityToDTO_OneQuestion(result.Question))
}

func convertDTOToEntity_QuestionCreate(reqDTO dto.QuestionCreate_Request) entity.QuestionCreateParams {
	return entity.QuestionCreateParams{
		Text: reqDTO.Text,
	}
}
