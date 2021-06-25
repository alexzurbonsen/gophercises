package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename_csv := flag.String("csv", "problems.csv", "a csv file providing questions and answers for the quiz in the form of 'question,answer'")
	flag.Parse()
	// fmt.Println("flags have been parsed: ", flag.Parsed())

	f, err := os.Open(*filename_csv)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file %s", *filename_csv))
	}
	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(questions)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) //trims preceeding and trailing spaces
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("Congratulations, you answered %.2f %% of the quiz correctly (%v / %v)\n", float64(correct)/float64(len(problems)), correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
