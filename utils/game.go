package utils

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"

	"github.com/chancehl/jeopardy-trainer/structs"
)

const CATEGORY_COUNT = 6

func GenerateSeed(rounds []structs.JeopardyRound) string {
	questions := []structs.JeopardyQuestion{}

	for _, value := range rounds {
		questions = append(questions, value.Questions...)
	}

	ids := make([]string, len(questions))

	for index, question := range questions {
		ids[index] = strconv.Itoa(question.Id)
	}

	seed := base64.StdEncoding.EncodeToString([]byte(strings.Join(ids, ",")))

	return seed
}

func GenerateRoundsFromSeed(seed string, allQuestions []structs.JeopardyQuestion) []structs.JeopardyRound {
	value, _ := base64.StdEncoding.DecodeString(seed)

	stringIds := strings.Split(string(value), ",")

	var ids []int

	for _, str := range stringIds {
		i, _ := strconv.Atoi(str)

		ids = append(ids, i)
	}

	jeopardyQuestions := []structs.JeopardyQuestion{}
	doubleJeopardyQuestions := []structs.JeopardyQuestion{}
	finalJeopardyQuestions := []structs.JeopardyQuestion{}

	for _, id := range ids {
		for _, question := range allQuestions {
			if id == question.Id {
				if question.Round == "Jeopardy" {
					jeopardyQuestions = append(jeopardyQuestions, question)
				} else if question.Round == "DoubleJeopardy" {
					doubleJeopardyQuestions = append(doubleJeopardyQuestions, question)
				} else {
					finalJeopardyQuestions = append(finalJeopardyQuestions, question)
				}
			}
		}
	}

	jeopardyRound := structs.JeopardyRound{
		Name:      "Jeopardy",
		Questions: jeopardyQuestions,
	}

	doubleJeopardyRound := structs.JeopardyRound{
		Name:      "DoubleJeopardy",
		Questions: doubleJeopardyQuestions,
	}
	
	finalJeopardyRound := structs.JeopardyRound{
		Name:      "FinalJeopardy",
		Questions: finalJeopardyQuestions,
	}

	return []structs.JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
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
		questions = append(questions, PickRandomQuestion(validQuestions))
	} else {
		categories := 0

		for categories < CATEGORY_COUNT {
			questions = append(questions, PickRandomQuestionCategory(validQuestions)...)

			categories++
		}
	}

	return questions
}

func GenerateRounds(allQuestions []structs.JeopardyQuestion) []structs.JeopardyRound {
	grouped := GroupByRound(allQuestions)

	jeopardyRound := structs.JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["Jeopardy"], false),
	}

	doubleJeopardyRound := structs.JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["DoubleJeopardy"], false),
	}

	finalJeopardyRound := structs.JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["FinalJeopardy"], true),
	}

	return []structs.JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
}

func GenerateJeopardyGame(allQuestions []structs.JeopardyQuestion) structs.JeopardyGame {
	var game structs.JeopardyGame

	game.Rounds = GenerateRounds(allQuestions)
	game.Seed = GenerateSeed(game.Rounds)

	return game
}
