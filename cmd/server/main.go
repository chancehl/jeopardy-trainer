package main

import (
	"github.com/chancehl/jeopardy-trainer/internal/handler"
	"github.com/chancehl/jeopardy-trainer/internal/hooks"
	"github.com/gin-gonic/gin"
)

func main() {
	// execute startup logic
	hooks.OnServerStart()

	router := gin.Default()

	// static assets
	router.Static("/assets", "./web/dist/assets")

	// ui routes
	router.GET("/", handler.ServeSPA)
	router.NoRoute(handler.HandleSPARoute)

	// game routes
	router.GET("/game/:seed", handler.GetGame)
	router.POST("/game", handler.CreateGame)

	// question routes
	router.POST("/questions/:id/validate", handler.ValidateAnswer)

	router.Run()
}
