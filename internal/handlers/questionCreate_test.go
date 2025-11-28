package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"question-answer-service/internal/dto"
	"question-answer-service/internal/entity"
	"question-answer-service/internal/enum"
	"question-answer-service/internal/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateQuestion(t *testing.T) {
	var (
		someValidQuestionText     = "some valid question text"
		failedValidationFieldText = "Key: 'QuestionCreate_Request.Text' Error:Field validation for 'Text' failed on the 'required' tag"
	)

	// Setup the mock repository and handler
	mockRepo := new(mocks.QuestionProvider)
	validator := validator.New(validator.WithRequiredStructEnabled())
	handler := NewQuestionHandler(mockRepo, validator)

	// Define test cases
	tests := []struct {
		name           string
		requestBody    string
		params         entity.QuestionCreateParams
		mockReturnData entity.QuestionCreateResult
		mockReturnErr  error
		expectedStatus int
		ErrorResponse  dto.ErrorResponse
	}{
		{
			name:        "Success case - create new question",
			requestBody: fmt.Sprintf(`{"text":"%v"}`, someValidQuestionText),
			params: entity.QuestionCreateParams{
				Text: someValidQuestionText,
			},
			mockReturnData: entity.QuestionCreateResult{
				Question: entity.Question{
					Id:        1,
					Text:      someValidQuestionText,
					CreatedAt: "2025-10-24T12:34:56Z",
				},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusCreated,
			ErrorResponse:  dto.ErrorResponse{},
		},
		{
			name:        "Bad request - empty request Body",
			requestBody: `{}`,
			mockReturnData: entity.QuestionCreateResult{
				Question: entity.Question{},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code:    enum.ERROR_RESPONSE_BAD_REQUEST,
				Message: failedValidationFieldText,
			},
		},
		{
			name:        "Bad request - wrong property name",
			requestBody: `{"wrong_name":"some text"}`,
			mockReturnData: entity.QuestionCreateResult{
				Question: entity.Question{},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code:    enum.ERROR_RESPONSE_BAD_REQUEST,
				Message: failedValidationFieldText,
			},
		},
		{
			name:        "Bad request - empty property value",
			requestBody: `{"text":""}`,
			mockReturnData: entity.QuestionCreateResult{
				Question: entity.Question{},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			ErrorResponse: dto.ErrorResponse{
				Code:    enum.ERROR_RESPONSE_BAD_REQUEST,
				Message: failedValidationFieldText,
			},
		},
		{
			name:        "Internal server error",
			requestBody: fmt.Sprintf(`{"text":"%v"}`, someValidQuestionText),
			params: entity.QuestionCreateParams{
				Text: someValidQuestionText,
			},
			mockReturnData: entity.QuestionCreateResult{},
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
			reqBodyReader := bytes.NewBufferString(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/questions", reqBodyReader)
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			// If we expect DB call
			if tt.expectedStatus == http.StatusCreated || tt.expectedStatus == http.StatusInternalServerError {
				// Mock the expected call for this specific test case
				mockRepo.On("QuestionCreate", mock.Anything, tt.params).
					Return(tt.mockReturnData, tt.mockReturnErr).Once()
			}

			// Call the handler function
			handler.CreateQuestion(rr, req)

			// Assertions using testify/assert
			assert.Equal(t, tt.expectedStatus, rr.Code, "Handler returned wrong status code")

			// Verify that the mock expectation was met
			mockRepo.AssertExpectations(t)

			// Check response body JSON structure,
			if tt.expectedStatus == http.StatusCreated {
				// Check successful response
				var responseBody dto.Question_Response
				err := json.NewDecoder(rr.Body).Decode(&responseBody)
				assert.Nil(t, err)
				assert.EqualValues(t,
					responseBody,
					convertEntityToDTO_OneQuestion(tt.mockReturnData.Question),
					"Wrong Successful response",
				)
			} else {
				// Check ErrorResponse correctness
				var errResponse dto.ErrorResponse
				err := json.NewDecoder(rr.Body).Decode(&errResponse)
				assert.Nil(t, err)
				assert.Equal(t, errResponse.Code, tt.ErrorResponse.Code)
				assert.EqualValues(t, errResponse, tt.ErrorResponse)
			}
		})
	}
}
