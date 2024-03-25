package main

import (
	"net/http"

	"github.com/chancehl/jeopardy-trainer/structs"
	"github.com/chancehl/jeopardy-trainer/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("./templates/index.html")
	})

	router.GET("/game/:seed", func(ctx *gin.Context) {
		allQuestions := utils.LoadQuestions("./questions.json")

		seed := ctx.Param("seed")
		questions := utils.GenerateQuestionsFromSeed(seed, allQuestions);

		game := structs.JeopardyGame {
			Seed: seed,
			Questions: questions,
		}

		ctx.JSON(http.StatusOK, gin.H{"game": game})
	})

	router.POST("/games", func(ctx *gin.Context) {
		questions := utils.LoadQuestions("./questions.json")

		game := utils.GenerateJeopardyGame(questions)

		ctx.JSON(http.StatusOK, gin.H{"game": game})
	})

	router.Run()
}
