# Jeopardy Trainer

## Overview

A web application that generates a Jeopardy game for you to play based off of past episode data.

## Technology

This project serves a Single Page Application (SPA) built with React (using Vite for bundling) through a Go server powered by the Gin framework.

## Prerequisites

Before setting up the project, ensure you have the following installed:

- **Go**: Version 1.16 or higher.
  - You can download it from [the official Go website](https://golang.org/dl/).
- **Node.js**: Version 14.x or higher, which includes `npm` for managing frontend dependencies.
  - Download Node.js from [the official Node.js website](https://nodejs.org/).
- **Git** (optional): For cloning the repository.
  - Download Git from [git-scm.com](https://git-scm.com/).

## Setup

### Seed the questions

First, clone the questions repository and then run the scripts inside to massage the data into the format needed for the Jeopardy Trainer app.

```bash
git clone https://github.com/chancehl/JeopardyQuestions

cd ./JeopardyQuestions

# grant permissions to scripts
chmod +x ./combine.sh
chmod +x ./format.js

# run scripts (the output of this should be a "combined.json" file)
./combine.sh && ./format.js
```

You should see some output that looks like this if the previous step was successful:

```
Combining files...
Processing file 1: ./src/episodes_7001_7500.json
Processing file 2: ./src/episodes_1001_1500.json
Processing file 3: ./src/episodes_2501_3000.json
Processing file 4: ./src/episodes_3501_4000.json
Processing file 5: ./src/episodes_4501_5000.json
Processing file 6: ./src/episodes_1501_2000.json
Processing file 7: ./src/episodes_501_1000.json
Processing file 8: ./src/episodes_001_500.json
Processing file 9: ./src/episodes_5501_6000.json
Processing file 10: ./src/episodes_8001_8500.json
Processing file 11: ./src/episodes_5001_5500.json
Processing file 12: ./src/episodes_6001_6500.json
Processing file 13: ./src/episodes_3001_3500.json
Processing file 14: ./src/episodes_8501_9000.json
Processing file 15: ./src/episodes_4001_4500.json
Processing file 16: ./src/episodes_6501_7000.json
Processing file 17: ./src/episodes_2001_2500.json
Successfully combined files into ./combined.json
Formatting combined.json...
Successfully formatted question data
```

### Clone the repository

Next, clone the Jeopardy Trainer repository to your local machine:

**Note**: This should not live inside of the JeopardyQuestions/ directory.

```bash
git clone https://github.com/chancehl/jeopardy-trainer.git

cd jeopardy-trainer
```

Lastly, copy over the `combined.json` file from the first step.

```bash
mv ../JeopardyQuestions/combined.json ../jeopardy-trainer/questions.json
```

### Frontend Setup

Navigate to the web directory and install the dependencies:

```bash
cd web
npm install
```

#### Build the frontend application:

```bash
npm run build
```

This command compiles your React application and outputs the static files to web/dist.

### Backend Setup

#### Install Go dependencies:

```bash
go mod tidy
```

#### Run the Go server:

```bash
go run cmd/server/main.go
```

The server will start, and you should see output indicating it's serving your SPA on a local port (default: :8080).

## Usage

After starting the server, open a web browser and navigate to http://localhost:8080 to view the SPA.

## Project Structure

- **`/cmd/server/main.go`**: The entry point for the Gin server.
- **`/internal`**: Contains the business logic for parsing Jeopardy questions and generating games.
  - **`/model`**: Domain types like `JeopardyGame` and `JeopardyQuestion`.
  - **`/parser`**: Logic for parsing the questions from `questions.json`.
- **`/web`**: Contains the React application source and build output.
  - **`/src`**: React source files.
  - **`/dist`**: Output from the `npm run build` command.
- **`questions.json`**: The dataset of Jeopardy questions.

This README was shamelessly [generated with LLM labor](https://github.com/chancehl/gpt-readme)💪🤖
