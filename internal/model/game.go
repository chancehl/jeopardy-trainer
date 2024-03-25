package model

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
)

const CATEGORY_COUNT = 6

func GenerateSeed(rounds []JeopardyRound) string {
	questions := []JeopardyQuestion{}

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

func GenerateRoundsFromSeed(seed string, allQuestions []JeopardyQuestion) []JeopardyRound {
	value, _ := base64.StdEncoding.DecodeString(seed)

	stringIds := strings.Split(string(value), ",")

	var ids []int

	for _, str := range stringIds {
		i, _ := strconv.Atoi(str)

		ids = append(ids, i)
	}

	jeopardyQuestions := []JeopardyQuestion{}
	doubleJeopardyQuestions := []JeopardyQuestion{}
	finalJeopardyQuestions := []JeopardyQuestion{}

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

	jeopardyRound := JeopardyRound{
		Name:      "Jeopardy",
		Questions: jeopardyQuestions,
	}

	doubleJeopardyRound := JeopardyRound{
		Name:      "DoubleJeopardy",
		Questions: doubleJeopardyQuestions,
	}

	finalJeopardyRound := JeopardyRound{
		Name:      "FinalJeopardy",
		Questions: finalJeopardyQuestions,
	}

	return []JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
}

func GroupByRound(questions []JeopardyQuestion) map[string][]JeopardyQuestion {
	grouped := make(map[string][]JeopardyQuestion)

	for _, question := range questions {
		grouped[question.Round] = append(grouped[question.Round], question)
	}

	return grouped
}

func PickRandomQuestionCategory(questions []JeopardyQuestion) []JeopardyQuestion {
	categoryQuestions := []JeopardyQuestion{}

	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	for _, question := range questions {
		if question.GameId == randomQuestion.GameId && question.Category == randomQuestion.Category {
			categoryQuestions = append(categoryQuestions, question)
		}
	}

	return categoryQuestions
}

func PickRandomQuestion(questions []JeopardyQuestion) JeopardyQuestion {
	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	return randomQuestion
}

func PickRandomQuestions(validQuestions []JeopardyQuestion, isFinalJeopardy bool) []JeopardyQuestion {
	questions := []JeopardyQuestion{}

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

func GenerateRounds(allQuestions []JeopardyQuestion) []JeopardyRound {
	grouped := GroupByRound(allQuestions)

	jeopardyRound := JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["Jeopardy"], false),
	}

	doubleJeopardyRound := JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["DoubleJeopardy"], false),
	}

	finalJeopardyRound := JeopardyRound{
		Name:      "Jeopardy",
		Questions: PickRandomQuestions(grouped["FinalJeopardy"], true),
	}

	return []JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
}

func GenerateJeopardyGame(allQuestions []JeopardyQuestion) JeopardyGame {
	var game JeopardyGame

	game.Rounds = GenerateRounds(allQuestions)
	game.Seed = GenerateSeed(game.Rounds)

	return game
}
