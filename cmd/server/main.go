package main

import (
	"net/http"

	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/chancehl/jeopardy-trainer/internal/parser"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/game/:seed", func(ctx *gin.Context) {
		allQuestions := parser.LoadQuestions("./questions.json")

		seed := ctx.Param("seed")
		rounds := model.GenerateRoundsFromSeed(seed, allQuestions)

		game := model.JeopardyGame{
			Seed:   seed,
			Rounds: rounds,
		}

		ctx.JSON(http.StatusOK, gin.H{"game": game})
	})

	router.POST("/games", func(ctx *gin.Context) {
		questions := parser.LoadQuestions("./questions.json")

		game := model.GenerateJeopardyGame(questions)

		ctx.JSON(http.StatusOK, gin.H{"game": game})
	})

	router.Run()
}
