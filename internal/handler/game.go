package handler

import (
	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/chancehl/jeopardy-trainer/internal/parser"
	"github.com/gin-gonic/gin"
)

func GetGame(ctx *gin.Context) {
	allQuestions := parser.LoadQuestions("./questions.json")

	seed := ctx.Param("seed")
	rounds := model.GenerateRoundsFromSeed(seed, allQuestions)

	game := model.JeopardyGame{
		Seed:   seed,
		Rounds: rounds,
	}

	body := gin.H{"game": game}

	ctx.JSON(200, body)
}

func CreateGame(ctx *gin.Context) {
	questions := parser.LoadQuestions("questions.json")

	game := model.GenerateJeopardyGame(questions)

	body := gin.H{"game": game}

	ctx.JSON(200, body)
}
