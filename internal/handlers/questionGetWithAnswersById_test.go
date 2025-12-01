package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
	"question-answer-service/internal/mocks"
	"question-answer-service/internal/routes"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetQuestionWithAnswersById(t *testing.T) {
	// Setup the mock repository and handler
	mockRepo := new(mocks.QuestionProvider)
	validator := validator.New(validator.WithRequiredStructEnabled())
	handler := NewQuestionHandler(mockRepo, validator)

	// Define test cases
	tests := []struct {
		name           string
		pathParam      string
		params         entity.QuestionGetWithAnswersByIdParams
		mockReturnData entity.QuestionGetWithAnswersByIdResult
		mockReturnErr  error
		expectedStatus int
		ErrorResponse  dto.ErrorResponse
	}{
		// OK 200 - get question with list of related answers
		{
			name:      "Success case - get question with empty list of related answers",
			pathParam: "1",
			params: entity.QuestionGetWithAnswersByIdParams{
				QuestionId: 1,
			},
			mockReturnData: entity.QuestionGetWithAnswersByIdResult{
				ShortAnswers:  []entity.AnswerShort{},
				FoundQuestion: true,
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			ErrorResponse:  dto.ErrorResponse{},
		},
		{
			name:      "Success case - get question with no-empty list of related answers",
			pathParam: "1",
			params: entity.QuestionGetWithAnswersByIdParams{
				QuestionId: 1,
			},
			mockReturnData: entity.QuestionGetWithAnswersByIdResult{
				ShortAnswers: []entity.AnswerShort{
					{AnswerBaseFields: entity.AnswerBaseFields{
						AnswerId:  1,
						UserId:    "u1",
						Text:      "some answer text",
						CreatedAt: "2025-10-25T12:34:56Z",
					}},
					{AnswerBaseFields: entity.AnswerBaseFields{
						AnswerId:  2,
						UserId:    "u2",
						Text:      "another answer text",
						CreatedAt: "2025-10-25T15:34:56Z",
					}},
				},
				FoundQuestion: true,
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			ErrorResponse:  dto.ErrorResponse{},
		},
		// Bad request 400 - path parameter: [empty, invalid]
		{
			name:           "Bad request - EMPTY path parameter",
			pathParam:      "",
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code: enum.ERROR_RESPONSE_BAD_REQUEST,
			},
		},
		{
			name:           "Bad request - INVALID (string) path parameter",
			pathParam:      "abc",
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code: enum.ERROR_RESPONSE_BAD_REQUEST,
			},
		},
		// Bad request 400 - bad validation: [no, less than 1, equal to zero]
		{
			name:           "Bad request - INVALID (int_negative) path parameter",
			pathParam:      "-1",
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code: enum.ERROR_RESPONSE_BAD_REQUEST,
			},
		},
		{
			name:           "Bad request - INVALID (int_zero) path parameter",
			pathParam:      "0",
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code: enum.ERROR_RESPONSE_BAD_REQUEST,
			},
		},
		// Not found 404 - no question with given Id
		{
			name:      "Not found - no question with Id",
			pathParam: "1",
			params: entity.QuestionGetWithAnswersByIdParams{
				QuestionId: 1,
			},
			mockReturnData: entity.QuestionGetWithAnswersByIdResult{
				ShortAnswers:  []entity.AnswerShort{},
				FoundQuestion: false,
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusNotFound,
			ErrorResponse: dto.ErrorResponse{
				Code:    enum.ERROR_RESPONSE_NOT_FOUND,
				Message: "Вопрос не найден",
			},
		},
		// Internal server error 500 - db error
		{
			name:      "Internal server error - db error",
			pathParam: "1",
			params: entity.QuestionGetWithAnswersByIdParams{
				QuestionId: 1,
			},
			mockReturnData: entity.QuestionGetWithAnswersByIdResult{},
			mockReturnErr:  ErrDBConnectionFailed,
			expectedStatus: http.StatusInternalServerError,
			ErrorResponse: dto.ErrorResponse{
				Code:    enum.ERROR_RESPONSE_INTERNAL_SERVER_ERROR,
				Message: ErrDBConnectionFailed.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup request and a recorder
			req := httptest.NewRequest(http.MethodGet, "/questions/"+tt.pathParam, nil)
			req.SetPathValue(routes.PathQuestionId, tt.pathParam)
			rr := httptest.NewRecorder()

			// If we expect DB call
			if tt.expectedStatus != http.StatusBadRequest {
				// Mock the expected call for this specific test case
				mockRepo.On("QuestionGetWithAnswersById", mock.Anything, tt.params).
					Return(tt.mockReturnData, tt.mockReturnErr).Once()
			}

			// Call the handler function
			handler.GetQuestionWithAnswersById(rr, req)

			// Assertions using testify/assert
			assert.Equal(t, tt.expectedStatus, rr.Code, "Handler returned wrong status code")

			// Verify that the mock expectation was met
			mockRepo.AssertExpectations(t)

			// Check response body JSON structure,
			if tt.expectedStatus == http.StatusOK {
				// Check successful response correctness
				var responseBody dto.QuestionGetById_Response
				err := json.NewDecoder(rr.Body).Decode(&responseBody)
				assert.Nil(t, err)
				assert.EqualValues(t,
					responseBody,
					dto.QuestionGetById_Response{
						Question:     convertEntityToDTO_OneQuestion(tt.mockReturnData.Question),
						ShortAnswers: convertEntityToDTO_ManyShortAnswers(tt.mockReturnData.ShortAnswers),
					},
					"Wrong Successful response",
				)
			} else {
				// Check ErrorResponse correctness
				var errResponse dto.ErrorResponse
				err := json.NewDecoder(rr.Body).Decode(&errResponse)
				assert.Nil(t, err)
				assert.Equal(t, errResponse.Code, tt.ErrorResponse.Code)
			}
		})
	}
}
