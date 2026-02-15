package main

import (
	"flag"
	"fmt"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "default csv file for quizzes, format: question,answer")
	flag.Parse()
	fmt.Println(*csvFileName)
}
