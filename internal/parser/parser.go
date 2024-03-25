package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chancehl/jeopardy-trainer/internal/model"
)

func loadJson[T any](filename string) (T, error) {
	var data T

	fileData, err := os.ReadFile(filename)

	if err != nil {
		return data, err
	}

	return data, json.Unmarshal(fileData, &data)
}

func LoadQuestions(filename string) []model.JeopardyQuestion {
	questions, err := loadJson[[]model.JeopardyQuestion](filename)

	if err != nil {
		fmt.Println(err)
	}

	return questions
}
