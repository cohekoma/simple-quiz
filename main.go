package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "quizzes.csv", "default csv file for quizzes, format: question,answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		quitApp(fmt.Sprintf("Failed to open CSV file: %s. File may not exist!", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		quitApp("Failed to parse the CSV file! Try using another file")
	}

	parsedQuizzes := parseRows(lines)
	playQuiz(parsedQuizzes)
}

type quiz struct {
	question string
	answer   string
}

func parseRows(lines [][]string) (quizzes []quiz) {
	result := make([]quiz, len(lines))
	for i, line := range lines {
		result[i] = quiz{
			question: line[0],
			answer:   line[1],
		}
	}

	return result
}

func playQuiz(quizzes []quiz) {
	score := 0
	for i, quiz := range quizzes {
		fmt.Printf("Question #%d: %s\n", i+1, quiz.question)
		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == quiz.answer {
			score++
		}
	}

	fmt.Printf("You got %d points out of %d questions!\n", score, len(quizzes))
}

func quitApp(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
