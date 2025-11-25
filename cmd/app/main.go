package main

import "question-answer-service/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}
