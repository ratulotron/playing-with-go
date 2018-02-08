package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Question struct {
	Question  string
	Answer    string
	IsCorrect bool
}

// check is an error handler
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadQuestions(filename string) (problems []Question) {
	file, err := os.Open(filename)
	check(err)
	// Close the file at the end of the program
	defer file.Close()

	// Read File into a variable
	lines, err := csv.NewReader(file).ReadAll()
	check(err)

	// Loop through lines & turn into object
	for _, line := range lines {
		question := Question{
			Question: line[0],
			Answer:   line[1],
		}
		problems = append(problems, question)
	}
	return problems
}

func askQuestion(question *Question) (isCorrect bool) {
	fmt.Print(question.Question, " = ")
	var answer string
	_, err := fmt.Scan(&answer)
	check(err)

	if question.Answer == answer {
		question.IsCorrect = true
	}
	return isCorrect
}

func main() {
	// flag, default, description
	filename := flag.String("word", "test.csv", "a string")
	// Unless this is called, command-line flags aren't really parsed
	flag.Parse()

	// Loads all the problems in the program
	questions := loadQuestions(*filename)

	result := 0
	for key, question := range questions {
		fmt.Printf("Problem #%d: ", key+1)
		if askQuestion(&question) {
			result++
		}
	}
	fmt.Println("Correctly answered: ", result)
}
