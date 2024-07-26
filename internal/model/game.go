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
		for _, category := range value.Categories {
			questions = append(questions, category.Questions...)
		}
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
		Name:       "Jeopardy",
		Categories: GenerateCategoriesFromQuestions(jeopardyQuestions),
	}

	doubleJeopardyRound := JeopardyRound{
		Name:       "DoubleJeopardy",
		Categories: GenerateCategoriesFromQuestions(doubleJeopardyQuestions),
	}

	finalJeopardyRound := JeopardyRound{
		Name:       "FinalJeopardy",
		Categories: GenerateCategoriesFromQuestions(finalJeopardyQuestions),
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

func GenerateCategoriesFromQuestions(questions []JeopardyQuestion) []JeopardyCategory {
	categoryMap := make(map[string][]JeopardyQuestion)

	for _, q := range questions {
		categoryMap[q.Category] = append(categoryMap[q.Category], q)
	}

	categories := make([]JeopardyCategory, 0, len(categoryMap))

	for name, qs := range categoryMap {
		categories = append(categories, JeopardyCategory{Name: name, Questions: qs})
	}

	return categories
}

func GenerateRounds(allQuestions []JeopardyQuestion) []JeopardyRound {
	grouped := GroupByRound(allQuestions)

	jeopardyRoundQuestions := PickRandomQuestions(grouped["Jeopardy"], false)
	doubleJeopardyRoundQuestions := PickRandomQuestions(grouped["DoubleJeopardy"], false)
	finalJeopardyRoundQuestions := PickRandomQuestions(grouped["FinalJeopardy"], true)

	jeopardyRound := JeopardyRound{
		Name:       "Jeopardy",
		Categories: GenerateCategoriesFromQuestions(jeopardyRoundQuestions),
	}

	doubleJeopardyRound := JeopardyRound{
		Name:       "DoubleJeopardy",
		Categories: GenerateCategoriesFromQuestions(doubleJeopardyRoundQuestions),
	}

	finalJeopardyRound := JeopardyRound{
		Name:       "FinalJeopardy",
		Categories: GenerateCategoriesFromQuestions(finalJeopardyRoundQuestions),
	}

	return []JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
}

func GenerateJeopardyGame(allQuestions []JeopardyQuestion) JeopardyGame {
	var game JeopardyGame

	game.Rounds = GenerateRounds(allQuestions)
	game.Seed = GenerateSeed(game.Rounds)

	return game
}
