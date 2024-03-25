package utils

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"

	"github.com/chancehl/jeopardy-trainer/structs"
)

const CATEGORY_COUNT = 6

func GenerateSeed(questions []structs.JeopardyQuestion) string {
    ids := make([]string, len(questions))

	for index, question := range questions {
		ids[index] = strconv.Itoa(question.Id)
	}

	seed := base64.StdEncoding.EncodeToString([]byte(strings.Join(ids, ",")))

	return seed
}

func GenerateQuestionsFromSeed(seed string, allQuestions []structs.JeopardyQuestion) []structs.JeopardyQuestion {
	value, _ := base64.StdEncoding.DecodeString(seed)

	stringIds := strings.Split(string(value), ",")

	var ids []int

	for _, str := range stringIds {
		i, _ := strconv.Atoi(str)

		ids = append(ids, i)
	}
	
	questions := []structs.JeopardyQuestion{}

	for _, id := range ids {
		for _, question := range allQuestions {
			if id == question.Id {
				questions = append(questions, question)
			}
		}
	}

	return questions
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