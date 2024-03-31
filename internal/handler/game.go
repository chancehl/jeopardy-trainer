package handler

import (
	"github.com/chancehl/jeopardy-trainer/internal/db"
	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/gin-gonic/gin"
)

func GetGame(ctx *gin.Context) {
	allQuestions := db.LoadQuestions("questions.json")

	seed := ctx.Param("seed")
	rounds := model.GenerateRoundsFromSeed(seed, allQuestions)

	game := model.JeopardyGame{
		Seed:   seed,
		Rounds: rounds,
	}

	ctx.JSON(200, game)
}

func CreateGame(ctx *gin.Context) {
	questions := db.LoadQuestions("questions.json")

	game := model.GenerateJeopardyGame(questions)

	ctx.JSON(200, game)
}
