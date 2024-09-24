import React from 'react'

import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { HomePage } from './pages'
import './App.css'

const queryClient = new QueryClient()

function App() {
    return (
        <React.Fragment>
            <QueryClientProvider client={queryClient}>
                <HomePage />
            </QueryClientProvider>
        </React.Fragment>
    )
}

export default App
