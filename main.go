package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JeopardyGame struct {
	Seed string
	Questions []JeopardyQuestion
}

type JeopardyQuestion struct {
	Question string
	Value int32
	Id string
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/", func(ctx *gin.Context) {
        game := JeopardyGame{
            Seed: "12345abcde",
			Questions: []JeopardyQuestion {
				{
					Question: "Who wrote 'To Kill a Mockingbird'?",
					Value: 200,
					Id: "1",
				},
				{
					Question: "Who wrote 'To Kill a Mockingbird'?",
					Value: 200,
					Id: "2",
				},
			},
        }
        
        ctx.HTML(200, "index.html", game)
	})

	router.Run()
}