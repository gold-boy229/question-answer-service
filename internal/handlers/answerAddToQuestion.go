package handlers

import (
	"encoding/json"
	"net/http"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
	"question-answer-service/internal/routes"
	"strconv"
)

func (h *answerHandler) AddAnswerToQuestion(w http.ResponseWriter, req *http.Request) {
	reqDTO, err := bindInputParams_AnswerAdd(req)
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

	params := convertDTOToEntity_AnsertCreate(reqDTO)
	result, err := h.repo.AnswerCreate(req.Context(), params)
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

	ResponseWithJSON(w, http.StatusCreated,
		dto.AnswerCreate_Response{
			Answer: convertEntityToDTO_OneAnswer(result.Answer),
		},
	)
}

func bindInputParams_AnswerAdd(req *http.Request) (dto.AnswerCreate_Request, error) {
	var reqDTO dto.AnswerCreate_Request
	defer req.Body.Close()

	err := json.NewDecoder(req.Body).Decode(&reqDTO)
	if err != nil {
		return dto.AnswerCreate_Request{}, err
	}

	reqDTO.QuestionId, err = getQuestionIdFromPath(req)
	if err != nil {
		return dto.AnswerCreate_Request{}, err
	}

	return reqDTO, nil
}

func getQuestionIdFromPath(req *http.Request) (*int, error) {
	questionIdStr := req.PathValue(routes.PathQuestionId)
	val, err := strconv.Atoi(questionIdStr)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func convertDTOToEntity_AnsertCreate(reqDTO dto.AnswerCreate_Request) entity.AnswerCreateParams {
	return entity.AnswerCreateParams{
		QuestionId: *reqDTO.QuestionId,
		UserId:     reqDTO.UserId,
		Text:       reqDTO.Text,
	}
}

func convertEntityToDTO_OneAnswer(ans entity.Answer) dto.Answer_Response {
	return dto.Answer_Response{
		AnswerId:   ans.AnswerId,
		QuestionId: ans.QuestionId,
		UserId:     ans.UserId,
		Text:       ans.Text,
		CreatedAt:  ans.CreatedAt,
	}
}
