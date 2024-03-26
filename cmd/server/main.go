package main

import (
	"fmt"

	"github.com/chancehl/jeopardy-trainer/internal/handler"
	"github.com/chancehl/jeopardy-trainer/internal/hooks"
	"github.com/gin-gonic/gin"
)

func main() {
	// execute startup logic
	hooks.OnServerStart()

	router := gin.Default()

	// static
	router.Static("/assets", "./web/dist/assets")

	// ui
	router.GET("/", handler.ServeSPA)
	router.NoRoute(handler.HandleSPARoute)

	// games
	router.GET("/games/:seed", handler.GetGame)
	router.POST("/games", handler.CreateGame)

	fmt.Println("Starting server...")
	router.Run()
}
