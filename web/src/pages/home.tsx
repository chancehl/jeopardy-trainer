import { useState } from 'react'
import { CopyIcon, ReloadIcon } from '@radix-ui/react-icons'
import { useQuery } from '@tanstack/react-query'

import { createGame } from '@/api'
import { Button } from '@/components'

export const HomePage = () => {
    const [enabled, setEnabled] = useState(false)

    const { data, isLoading } = useQuery({ queryKey: ['game'], queryFn: createGame, enabled })

    const onCopySeedClick = async () => {
        if (data?.game) {
            await navigator.clipboard.writeText(data?.game.seed)
        }
    }

    return (
        <>
            {data?.game == null && (
                <div className="flex items-center align-center gap-2">
                    <Button onClick={() => setEnabled(true)} disabled={isLoading}>
                        {isLoading && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}Play game
                    </Button>
                </div>
            )}
            {data?.game && data?.game?.seed && (
                <Button onClick={onCopySeedClick}>
                    <CopyIcon className="mr-2 h-4 w-4" />
                    Copy seed
                </Button>
            )}
        </>
    )
}
