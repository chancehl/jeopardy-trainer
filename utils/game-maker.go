package utils

import (
	"encoding/base64"
	"math/rand"
	"strings"

	"github.com/chancehl/jeopardy-trainer/structs"
)

const CATEGORY_COUNT = 6

func GenerateSeed(questions []structs.JeopardyQuestion) string {
	var ids []string

	for _, question := range questions {
		ids = append(ids, question.Id)
	}

	// Concatenate all prompts into a single string
	allPrompts := strings.Join(ids, "")

	// Encode the concatenated string using base64
	seed := base64.StdEncoding.EncodeToString([]byte(allPrompts))

	return seed
}

func GroupByRound(questions []structs.JeopardyQuestion) map[string][]structs.JeopardyQuestion {
	grouped := make(map[string][]structs.JeopardyQuestion)

	for _, question := range questions {
		grouped[question.Round] = append(grouped[question.Round], question)
	}

	return grouped
}

func PickRandomQuestionCategory(questions []structs.JeopardyQuestion) []structs.JeopardyQuestion {
	categoryQuestions := []structs.JeopardyQuestion{}

	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	for _, question := range questions {
		if question.GameId == randomQuestion.GameId && question.Category == randomQuestion.Category {
			categoryQuestions = append(categoryQuestions, question)
		}
	}

	return categoryQuestions
}

func PickRandomQuestion(questions []structs.JeopardyQuestion) structs.JeopardyQuestion {
	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	return randomQuestion
}

func PickRandomQuestions(validQuestions []structs.JeopardyQuestion, isFinalJeopardy bool) []structs.JeopardyQuestion {
	questions := []structs.JeopardyQuestion{}

	if isFinalJeopardy {
		questions = append(questions,PickRandomQuestion(validQuestions))
	} else {
		categories := 0
	
		for ; categories < CATEGORY_COUNT; {
			questions = append(questions, PickRandomQuestionCategory(validQuestions)...)
	
			categories++
		}
	}

	return questions
}

func GenerateQuestions(allQuestions []structs.JeopardyQuestion) []structs.JeopardyQuestion {
	questions := []structs.JeopardyQuestion{}
	
	grouped := GroupByRound(allQuestions)

	questions = append(questions, PickRandomQuestions(grouped["Jeopardy"], false)...)
	questions = append(questions, PickRandomQuestions(grouped["DoubleJeopardy"], false)...)
	questions = append(questions, PickRandomQuestions(grouped["FinalJeopardy"], true)...)

	return questions
}

func GenerateJeopardyGame(allQuestions []structs.JeopardyQuestion) structs.JeopardyGame {
	var game structs.JeopardyGame

	game.Questions = GenerateQuestions(allQuestions)
	game.Seed = GenerateSeed(game.Questions)

	return game
}