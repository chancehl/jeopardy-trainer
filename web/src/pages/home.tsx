import { CopyIcon, ReloadIcon } from '@radix-ui/react-icons'
import { useQuery } from '@tanstack/react-query'

import { createGame } from '@/api'
import { Button } from '@/components'
import { JeopardyGame } from '@/types'

export const HomePage = () => {
    const { data, isLoading, refetch } = useQuery<JeopardyGame>({ queryKey: ['game'], queryFn: createGame, enabled: false })

    const onCopySeedClick = async () => {
        if (data?.seed) {
            await navigator.clipboard.writeText(data?.seed)
        }
    }

    return (
        <>
            {data == null && (
                <div className="flex items-center align-center gap-2">
                    <Button onClick={() => refetch()} disabled={isLoading}>
                        {isLoading && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}Play game
                    </Button>
                </div>
            )}
            {data?.rounds.map((round) => (
                <div>
                    <h2>{round.name}</h2>
                </div>
            ))}
            {data?.seed && (
                <Button onClick={onCopySeedClick}>
                    <CopyIcon className="mr-2 h-4 w-4" />
                    Copy seed
                </Button>
            )}
        </>
    )
}
