package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var filename_csv string

func init() {
	flag.StringVar(&filename_csv, "csv", "problems.csv", "a csv file providing questions and answers for the quiz in the form of 'question,answer'")
}

func main() {

	flag.Parse()
	fmt.Println("flags have been parsed: ", flag.Parsed())
	f, err := os.Open(filename_csv)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	// declare count variable
	var count = 0

	for _, qa := range questions {
		fmt.Printf("%v = ", qa[0])
		answer, _ := reader.ReadString('\n')
		// convert CRLF to LF (CR: carriage return \r -> moves cursor to beginning of the line)
		// (LF: line feed \n -> moves cursor to next line without returning to the beginning)
		answer = strings.Replace(answer, "\n", "", -1)

		if strings.Compare(qa[1], answer) == 0 {
			count++
		}

	}
	fmt.Printf("Congratulations, you answered %v %% of the quiz correctly (%v / 12)\n", count/len(questions), count)

}
