import { useState } from 'react'
import { ReloadIcon } from '@radix-ui/react-icons'

import { Button } from '@/components/ui/button'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import { Textarea } from '@/components/ui/textarea'

type Props = {
    onSubmit: (seed: string) => Promise<void>
    loading?: boolean
}

export function LoadFromDialog({ onSubmit, loading = false }: Props) {
    const [seed, setSeed] = useState('')

    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button variant="outline">Load game</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Load game from seed</DialogTitle>
                    <DialogDescription>This will allow you to load a past game based on a saved seed.</DialogDescription>
                </DialogHeader>
                <div className="flex flex-col gap-4 py-4">
                    <Textarea id="name" value={seed} className="col-span-3" onChange={(e) => setSeed(e.target.value)} />
                </div>
                <DialogFooter>
                    <Button type="submit" disabled={seed.length === 0 || loading} onClick={() => onSubmit(seed)}>
                        {loading && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}
                        Load from seed
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}
