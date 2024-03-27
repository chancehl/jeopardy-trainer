export type JeopardyQuestion = {
    prompt: string
    category: string
    round: string
    value: number
    answer: string
    id: number
    gameId: number
}

export type JeopardyRound = {
    name: string
    questions: JeopardyQuestion[]
}

export type JeopardyGame = {
    seed: string
    rounds: JeopardyRound[]
}
