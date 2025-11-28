package handlers

import (
	"encoding/json"
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

func TestGetAllQuestions(t *testing.T) {
	// Setup the mock repository and handler
	mockRepo := new(mocks.QuestionProvider)
	validator := validator.New(validator.WithRequiredStructEnabled())
	handler := NewQuestionHandler(mockRepo, validator)

	// Define test cases
	tests := []struct {
		name           string
		mockReturnData []entity.Question
		mockReturnErr  error
		expectedStatus int
		ErrorResponse  dto.ErrorResponse
	}{
		{
			name:           "Success case - returns empty list of questions",
			mockReturnData: []entity.Question{},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			ErrorResponse:  dto.ErrorResponse{},
		},
		{
			name: "Success case - returns list with two questions",
			mockReturnData: []entity.Question{
				{Id: 1, Text: "Q1", CreatedAt: "2023-01-01 10:00:00"},
				{Id: 2, Text: "Q2", CreatedAt: "2023-01-01 11:00:00"},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			ErrorResponse:  dto.ErrorResponse{},
		},
		{
			name:           "Internal Server Error case",
			mockReturnData: nil,
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
			// Mock the expected call for this specific test case
			mockRepo.On("QuestionGetAll", mock.Anything).
				Return(tt.mockReturnData, tt.mockReturnErr).Once()

			// Create a request and a recorder to capture the response
			req := httptest.NewRequest(http.MethodGet, "/questions", nil)
			rr := httptest.NewRecorder()

			// Call the handler function
			handler.GetAllQuestions(rr, req)

			// Assertions using testify/assert
			assert.Equal(t, tt.expectedStatus, rr.Code, "Handler returned wrong status code")

			// Verify that the mock expectation was met
			mockRepo.AssertExpectations(t)

			// Check response body JSON structure,
			if tt.expectedStatus == http.StatusOK {
				// Check successful response
				var responseBody dto.QuestionGetAll_Response
				err := json.NewDecoder(rr.Body).Decode(&responseBody)
				assert.Nil(t, err)
				assert.EqualValues(t,
					responseBody,
					dto.QuestionGetAll_Response{
						Questions: convertEntityToDTO_ManyQuestions(tt.mockReturnData),
					},
					"Wrong Successful response",
				)
			} else {
				// Check ErrorResponse correctness
				var errResponse dto.ErrorResponse
				err := json.NewDecoder(rr.Body).Decode(&errResponse)
				assert.Nil(t, err)
				assert.EqualValues(t, errResponse, tt.ErrorResponse, "Wrong ErrorResponse")
			}
		})
	}
}
