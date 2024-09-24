import client from './client'

export const createGame = async () => {
    const response = await client.post('/games')

    return response.data
}
