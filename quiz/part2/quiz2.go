package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var filename_csv string
var t int

func init() {
	flag.StringVar(&filename_csv, "csv", "problems.csv", "a csv file providing questions and answers for the quiz in the form of 'question,answer'")
	flag.IntVar(&t, "time", 2, "the time limit for each answer in seconds (integer)")
}

func main() {

	flag.Parse()
	if !flag.Parsed() {
		fmt.Println("Error: flags have not been parsed.")
		os.Exit(1)
	}

	f, err := os.Open(filename_csv)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var count float64 = 0
	var i = 0
	c := make(chan float64)

	for i = 0; i < len(questions); i++ {
		go quiz(questions, i, c)

		select {
		case res := <-c:
			count += res
		case <-time.After(time.Duration(t) * time.Second):
			fmt.Println("time passed")
			break // this is not necessary but may be better style?
		}
	}

	fmt.Printf("\nCongratulations, you answered %.2f %% of the quiz correctly (%v / 12)\n", count/float64(len(questions)), count)
}

func quiz(qa [][]string, i int, c chan float64) {
	fmt.Printf("%v = ", qa[i][0])

	var answer string
	fmt.Scanf("%s", &answer)
	// if err != nil {
	// 	log.Fatal(err)
	// } else if n > 1 {
	// 	fmt.Println("ambigous answer given")
	// }

	if strings.Compare(qa[i][1], answer) == 0 {
		c <- 1
	} else {
		c <- 0
	}
}
