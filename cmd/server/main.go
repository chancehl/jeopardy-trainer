package main

import (
	"github.com/chancehl/jeopardy-trainer/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// static
	router.Static("/assets", "./web/dist/assets")

	// ui
	router.GET("/", handler.ServeSPA)
	router.NoRoute(handler.HandleSPARoute)

	// games
	router.GET("/games/:seed", handler.GetGame)
	router.POST("/games", handler.CreateGame)

	router.Run()
}
