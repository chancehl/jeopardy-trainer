package main

import (
	"github.com/chancehl/jeopardy-trainer/internal/endpoints"
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
	router.GET("/", endpoints.ServeSPA)
	router.NoRoute(endpoints.HandleSPARoute)

	// game routes
	router.GET("/game/:seed", endpoints.GetGame)
	router.POST("/game", endpoints.CreateGame)

	// question routes
	router.POST("/questions/:id/validate", endpoints.ValidateAnswer)

	router.Run()
}
