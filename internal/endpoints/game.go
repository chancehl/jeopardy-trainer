package endpoints

import (
	"github.com/chancehl/jeopardy-trainer/internal/db"
	"github.com/chancehl/jeopardy-trainer/internal/errors"
	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/chancehl/jeopardy-trainer/internal/service"
	"github.com/gin-gonic/gin"
)

func GetGame(ctx *gin.Context) {
	allQuestions, err := db.LoadQuestions("questions.json")
	if err != nil {
		ctx.JSON(500, errors.NewInternalServiceErrorJSON("failed to load questions", err))
		return
	}

	seed := ctx.Param("seed")
	rounds := service.GenerateRoundsFromSeed(seed, allQuestions)

	game := model.JeopardyGame{Seed: seed, Rounds: rounds}

	ctx.JSON(200, game)
}

func CreateGame(ctx *gin.Context) {
	questions, err := db.LoadQuestions("questions.json")
	if err != nil {
		ctx.JSON(500, errors.NewInternalServiceErrorJSON("could not load questions", err))
		return
	}

	game := service.GenerateJeopardyGame(questions)

	ctx.JSON(200, game)
}
