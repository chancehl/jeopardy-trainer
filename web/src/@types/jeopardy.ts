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
    categories: JeopardyCategory[]
}

export type JeopardyGame = {
    seed: string
    rounds: JeopardyRound[]
}

export type JeopardyCategory = {
    name: string
    questions: JeopardyQuestion[]
}
