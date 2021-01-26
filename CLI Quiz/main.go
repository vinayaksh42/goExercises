package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// function for reading a csv file
func readCsv(CsvFile string) [][]string {
	f, err := os.Open(CsvFile)
	if err != nil {
		log.Fatal("Unable to read the csv file: " + CsvFile)
	}
	// defers the execution of a function until the surrounding function returns
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+CsvFile, err)
	}

	return records
}

func main() {
	csvfile := "problems.csv"
	problems := readCsv(csvfile)

	// iterating over questions
	fmt.Println("Reading the current quiz from: ", csvfile)
	correct := 0

	timeLimit := flag.Int("Limit", 30, "The limit for the quiz in seconds")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for index, question := range problems {
		fmt.Print("(", index+1, ")", question[0], "?", ":")
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time is up")
			fmt.Println("You scored", correct, "out of", len(problems))
			return
		case answer := <-answerCh:
			if answer == question[1] {
				correct++
			}
		}
	}
	fmt.Println("You scored", correct, "out of", len(problems))
}
