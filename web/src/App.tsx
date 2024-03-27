import React, { useState } from 'react'
import { ReloadIcon, CopyIcon } from '@radix-ui/react-icons'

import './App.css'
import { Button } from './components/ui/button'
import { LoadFromDialog } from './components/ui/loadFromSeedDialog'
import { JeopardyGame } from './@types'

function App() {
    const [game, setGame] = useState<JeopardyGame>()
    const [loadingGame, setLoadingGame] = useState(false)
    const [loadingFromSeed, setLoadingFromSeed] = useState(false)

    const onPlayClick = async () => {
        try {
            setLoadingGame(true)

            const response = await fetch('/games', { method: 'POST' })

            const data = await response.json()

            setGame(data)
        } catch (err) {
            setLoadingGame(false)
        } finally {
            setLoadingGame(false)
        }
    }

    const onLoadFromSeedSubmit = async (seed: string) => {
        try {
            setLoadingFromSeed(true)

            const response = await fetch(`/games/${seed}`)

            const data = await response.json()

            setGame(data)
        } catch (err) {
            setLoadingFromSeed(false)
        } finally {
            setLoadingFromSeed(false)
        }
    }

    const onCopySeedClick = async () => {
        if (game) {
            await navigator.clipboard.writeText(game.seed)
        }
    }

    return (
        <React.Fragment>
            {game == null && (
                <div className="flex items-center align-center gap-2">
                    <Button onClick={onPlayClick} disabled={loadingGame}>
                        {loadingGame && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}Play game
                    </Button>
                    <LoadFromDialog onSubmit={onLoadFromSeedSubmit} loading={loadingFromSeed} />
                </div>
            )}
            {game && game.seed && (
                <Button onClick={onCopySeedClick}>
                    <CopyIcon className="mr-2 h-4 w-4" />
                    Copy seed
                </Button>
            )}
        </React.Fragment>
    )
}

export default App
