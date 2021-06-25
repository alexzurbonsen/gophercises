package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	filename_csv := flag.String("csv", "problems.csv", "a csv file providing questions and answers for the quiz in the form of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for an answer in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) //trims preceeding and trailing spaces, \n is added to catch the cases in which user presses enter after typing the answer
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime is up! You answered %.2f %% of the quiz correctly (%v / %v)\n", float64(correct)/float64(len(problems)), correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
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
