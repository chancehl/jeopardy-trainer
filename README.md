# Jeopardy Trainer

## Overview

This project serves a Single Page Application (SPA) built with React (using Vite for bundling) through a Go server powered by the Gin framework. The Go server is responsible for serving the static files generated by the React build process and providing API endpoints to interact with Jeopardy-style game data. The game logic includes parsing Jeopardy questions from a JSON file and generating random game rounds.

## Prerequisites

Before setting up the project, ensure you have the following installed:

- **Go**: Version 1.16 or higher. You can download it from [the official Go website](https://golang.org/dl/).
- **Node.js**: Version 14.x or higher, which includes `npm` for managing frontend dependencies. Download Node.js from [the official Node.js website](https://nodejs.org/).
- **Git** (optional): For cloning the repository. Download Git from [git-scm.com](https://git-scm.com/).

## Setup

### Clone the Repository

First, clone the repository to your local machine (skip this step if you have the project files already):

```bash
git clone https://github.com/yourusername/yourprojectname.git
cd yourprojectname
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
Copy code
go mod tidy
```

#### Run the Go server:

```bash
Copy code
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

> This README was shamelessly [generated with LLM labor 💪🤖](https://github.com/chancehl/gpt-readme)
