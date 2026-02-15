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

	fmt.Println(lines)
}

func quitApp(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
