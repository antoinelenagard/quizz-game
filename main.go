package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "quizz format is 'question,anwser'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV File")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)

	correct_counter := 0
	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = \n", i+1, p.question)
		var a string
		fmt.Scanf("%s\n", &a)

		if a == p.answer {
			correct_counter++
		}
	}
	fmt.Printf("You scored %d out of %d. \n", correct_counter, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
