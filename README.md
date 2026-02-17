# Super, Simple Quizzes

Simple quiz CLI application Written in Go. Quizzes can be prepared with a CSV file or fetched from an API.

## Requirement
- Go 1.24

## Usage

- Clone the repository then `cd` into it.
- Build with `go build .` (or just run it with `go run .`).

## Flags
- `csv`: specify your custom CSV file. If no CSV file is specify, the default sample `quizzes.csv` file will be parsed and used.
- `timeLimit`: time limit for each question, specified in seconds. If no time limit specified, default will be set to 2 seconds.