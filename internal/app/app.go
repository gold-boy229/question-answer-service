package app

import (
	"fmt"
	"net/http"
	"question-answer-service/internal/config"
	"question-answer-service/internal/handlers"
	"question-answer-service/internal/repository"
	"question-answer-service/internal/routes"

	"github.com/go-playground/validator/v10"
)

type App struct {
	server *http.Server
}

func NewApp() *App {
	configDB, err := config.ReadConfigDB()
	if err != nil {
		panic(fmt.Sprintf("Cannot load configDB: %v", err.Error()))
	}

	repo, err := repository.NewRepository(configDB)
	if err != nil {
		panic(fmt.Sprintf("Cannot create repository: %v", err.Error()))
	}

	validator := validator.New(validator.WithRequiredStructEnabled())

	var (
		mux                             = http.NewServeMux()
		questionHandler questionHandler = handlers.NewQuestionHandler(repo, validator)
		answerHandler   answerHandler   = handlers.NewAnswerHandler(repo, validator)
	)

	mux.HandleFunc(http.MethodGet+" /questions", questionHandler.GetAllQuestions)
	mux.HandleFunc(http.MethodPost+" /questions", questionHandler.CreateQuestion)
	mux.HandleFunc(http.MethodGet+" /questions/{"+routes.PathQuestionId+"}", questionHandler.GetQuestionWithAnswersById)
	mux.HandleFunc(http.MethodDelete+" /questions/{"+routes.PathQuestionId+"}", questionHandler.DeleteQuestionById)

	mux.HandleFunc(http.MethodPost+" /questions/{"+routes.PathQuestionId+"}/answers", answerHandler.AddAnswerToQuestion)
	mux.HandleFunc(http.MethodGet+" /answers/{"+routes.PathAnswerId+"}", answerHandler.GetAnswerById)
	mux.HandleFunc(http.MethodDelete+" /answers/{"+routes.PathAnswerId+"}", answerHandler.DeleteAnswerById)

	return &App{
		server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (a *App) Run() {
	if err := a.server.ListenAndServe(); err != nil {
		fmt.Printf("http.ListenAndServe err = %v\n", err)
	}
}
