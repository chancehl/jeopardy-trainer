import React, { useState } from 'react'
import { ReloadIcon, CopyIcon } from '@radix-ui/react-icons'

import './App.css'
import { Button } from './components/ui/button'

function App() {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const [game, setGame] = useState<any>(null)
    const [loadingGame, setLoadingGame] = useState(false)

    const onPlayClick = async () => {
        try {
            setLoadingGame(true)

            const response = await fetch('/games', { method: 'POST' })

            const data = await response.json()

            setGame(data)
        } catch (err) {
            setLoadingGame(false)
        }
    }

    const onCopySeedClick = async () => {
        await navigator.clipboard.writeText(game.seed)
    }

    return (
        <React.Fragment>
            {game == null && (
                <Button onClick={onPlayClick} disabled={loadingGame}>
                    {loadingGame && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}Play game
                </Button>
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
