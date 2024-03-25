package model

type JeopardyGame struct {
	Seed   string
	Rounds []JeopardyRound
}

type JeopardyQuestion struct {
	Prompt   string
	Category string
	Round    string
	Value    int
	Answer   string
	Id       int
	GameId   int
}

type JeopardyRound struct {
	Name      string
	Questions []JeopardyQuestion
}
