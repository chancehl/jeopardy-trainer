package utils

import (
	"encoding/base64"
	"fmt"
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

func PickRandomQuestionCategory(questions []structs.JeopardyQuestion) string {
	categories := []string{}

	for _, question := range questions {
		categories = append(categories, question.Category)
	}

	randomIndex := rand.Intn(len(categories))
	randomCategory := categories[randomIndex]

	return randomCategory
}

func PickQuestionsMatchingCategory(category string) []structs.JeopardyQuestion {
	questions := []structs.JeopardyQuestion{}

	for _, question := range questions {
		if question.Category == category {
			questions = append(questions, question)
		}
	}

	return questions
}

func PickRandomQuestions(round string, validQuestions []structs.JeopardyQuestion) []structs.JeopardyQuestion {
	questions := []structs.JeopardyQuestion{}

	categories := 6

	if round == "FinalJeopardy" {
		categories = 1
	}
	
	for ; categories >= 0; {
		randomCategory := PickRandomQuestionCategory(validQuestions)
		matchingQuestions := PickQuestionsMatchingCategory(randomCategory)

		questions = append(questions, matchingQuestions...)

		categories--
	}


	return questions
}

func GenerateQuestions(allQuestions []structs.JeopardyQuestion) []structs.JeopardyQuestion {
	questions := []structs.JeopardyQuestion{}
	
	grouped := GroupByRound(allQuestions)

	for key := range grouped {
        fmt.Println("Key:", key)
	}

	questions = append(questions, PickRandomQuestions("Jeopardy", grouped["Jeopardy"])...)
	questions = append(questions, PickRandomQuestions("DoubleJeopardy", grouped["DoubleJeopardy"])...)
	questions = append(questions, PickRandomQuestions("FinalJeopardy", grouped["FinalJeopardy"])...)

	return questions
}

func GenerateJeopardyGame(allQuestions []structs.JeopardyQuestion) structs.JeopardyGame {
	var game structs.JeopardyGame

	game.Questions = GenerateQuestions(allQuestions)
	game.Seed = GenerateSeed(game.Questions)

	return game
}