package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//flag variables
	limitFlag := flag.Int("limit", 30, "time to answer")
	csvFlag := flag.String("csv", "problems.csv", "csv file with questions")

	flag.Parse()

	limit := *limitFlag
	filename := *csvFlag

	timer := time.NewTimer(time.Duration(limit) * time.Second)
	timer.Stop()
	go func() {
		<-timer.C
		fmt.Println("Time out, you failed")
		os.Exit(0)
	}()

	fmt.Printf("limit = %v, filename = %v\n", limit, filename)

	data, err := os.ReadFile(filename)
	check(err)

	userInput := bufio.NewReader(os.Stdin)
	var answer string
	var correctAnswers int
	var totalAnswers int

	r := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		fmt.Print(record[0], "= ")

		timer.Reset(time.Duration(limit) * time.Second)
		answer, _ = userInput.ReadString('\n')
		answer = strings.TrimSpace(answer)

		totalAnswers += 1
		if answer == record[1] {
			correctAnswers += 1
		}
	}
	fmt.Println("You got", correctAnswers, "out of", totalAnswers)

}
