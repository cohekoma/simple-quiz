package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "quizzes.csv", "default csv file for quizzes, format: question,answer")
	timeLimit := flag.Int("timeLimit", 2, "time limit for each question, default is 2s.")
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
	playQuiz(parsedQuizzes, timeLimit)
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
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return result
}

func playQuiz(quizzes []quiz, timeLimit *int) {
	score := 0
	answerCh := make(chan string)
	for i, quiz := range quizzes {
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Question #%d: %s\n", i+1, quiz.question)

		go func() {
			userAnswer, err := reader.ReadString('\n')

			if err != nil {
				quitApp("Fail to read your input, try again.")
			}

			userAnswer = strings.TrimSpace(userAnswer)
			fmt.Printf("Your answer: %s\n", userAnswer)

			// var userAnswer string
			// fmt.Scanf("%s\n", &userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case ans := <-answerCh:
			if ans == quiz.answer {
				score++
			}
		case <-timer.C:
			quitApp("Timeout!")
		}
	}

	fmt.Printf("You got %d points out of %d questions!\n", score, len(quizzes))
}

func quitApp(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
