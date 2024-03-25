package main

import (
	"github.com/chancehl/jeopardy-trainer/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/games/:seed", handler.GetGame)
	router.POST("/games", handler.CreateGame)

	router.Run()
}
