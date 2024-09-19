package endpoints

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
		ctx.JSON(400, err.Error())
		return
	}

	questionId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, "invalid question id")
		return
	}

	var question model.JeopardyQuestion

	if err := db.FindQuestionById(&question, questionId); err != nil {
		distance, maxDistance := utils.Levenshtein(question.Answer, req.Answer)

		ctx.JSON(200, gin.H{
			"isCorrect":     distance < maxDistance,
			"correctAnswer": question.Answer,
			"distance":      distance,
			"maxDistance":   maxDistance,
		})
	} else {
		ctx.JSON(404, "Missing question")
	}
}
