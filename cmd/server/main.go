package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/chancehl/jeopardy-trainer/internal/handler"
	"github.com/gin-gonic/gin"
)

func buildWebPackage() {
	fmt.Println("Building web package...")

	cmd := exec.Command("npm", "run", "build")

	cmd.Dir = "./web"

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Failed to build web package: %s\n", err)
	}

	fmt.Printf("[npm run build]: %s\n", output)
}

func main() {
	// run this on startup
	buildWebPackage()

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
