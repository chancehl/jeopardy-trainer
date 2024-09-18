package handler

import (
	"fmt"

	"github.com/chancehl/jeopardy-trainer/internal/db"
	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/gin-gonic/gin"
)

func GetGame(ctx *gin.Context) {
	allQuestions, err := db.LoadQuestions("questions.json")
	if err != nil {
		fmt.Printf("failed to load questions from json: %v\n", err)
		return
	}

	seed := ctx.Param("seed")
	rounds := model.GenerateRoundsFromSeed(seed, allQuestions)

	ctx.JSON(200, model.JeopardyGame{Seed: seed, Rounds: rounds})
}

func CreateGame(ctx *gin.Context) {
	questions, err := db.LoadQuestions("questions.json")
	if err != nil {
		fmt.Printf("could not load questions from json: %v\n", err)
		return
	}

	ctx.JSON(200, model.GenerateJeopardyGame(questions))
}
