package model

type JeopardyGame struct {
	Seed   string          `json:"seed"`
	Rounds []JeopardyRound `json:"rounds"`
}

type JeopardyQuestion struct {
	Prompt   string `json:"prompt"`
	Category string `json:"category"`
	Round    string `json:"round"`
	Value    int    `json:"value"`
	Answer   string `json:"answer"`
	Id       int    `json:"id"`
	GameId   int    `json:"gameId"`
}

type JeopardyRound struct {
	Name       string             `json:"name"`
	Categories []JeopardyCategory `json:"categories"`
}

type JeopardyCategory struct {
	Name      string             `json:"name"`
	Questions []JeopardyQuestion `json:"questions"`
}
