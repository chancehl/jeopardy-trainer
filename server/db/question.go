package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chancehl/jeopardy-trainer/server/model"
)

func loadJson[T any](filename string) (T, error) {
	var data T

	fileData, err := os.ReadFile(filename)

	if err != nil {
		return data, err
	}

	return data, json.Unmarshal(fileData, &data)
}

func LoadQuestions(filename string) ([]model.JeopardyQuestion, error) {
	questions, err := loadJson[[]model.JeopardyQuestion](filename)

	if err != nil {
		return nil, fmt.Errorf("could not load questions from file: %v", err)
	}

	return questions, nil
}

func FindQuestionById(question *model.JeopardyQuestion, id int) error {
	questions, err := LoadQuestions("questions.json")
	if err != nil {
		return fmt.Errorf("error loading questions from json: %v", err)
	}

	for _, q := range questions {
		if q.Id == id {
			*question = q
			return nil
		}
	}

	return fmt.Errorf("could not locate question by id %d", id)
}
