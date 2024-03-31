package handler

import (
	"strconv"

	"github.com/chancehl/jeopardy-trainer/internal/db"
	"github.com/chancehl/jeopardy-trainer/internal/model"
	"github.com/chancehl/jeopardy-trainer/internal/utils"
	"github.com/gin-gonic/gin"
)

func ValidateAnswer(ctx *gin.Context) {
	var req model.UserAnswer

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})

		return
	}

	questionId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})

		return
	}

	question, found := db.GetQuestionById(questionId)

	if found {
		distance, maxDistance := utils.Levenshtein(question.Answer, req.Answer)

		ctx.JSON(200, gin.H{
			"correct":       distance < maxDistance,
			"correctAnswer": question.Answer,
			"dist":          distance,
			"max":           maxDistance,
		})
	} else {
		ctx.JSON(404, "Missing question")
	}
}
