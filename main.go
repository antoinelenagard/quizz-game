package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "quizz format is 'question,anwser'")
	flag.Parse()
	// _ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: ", *csvFilename)
		os.Exit(1)
	}

	_ = file
}
