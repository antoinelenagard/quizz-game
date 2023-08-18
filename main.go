package main 

import "flag"

func main () {
	csvFilename := flag.String("csv", "problem.csv", "quizz format is 'question,anwser")
	flag.Parse()
	_ = csvFilename
}
