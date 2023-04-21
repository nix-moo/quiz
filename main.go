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
	fileName := flag.String("csv", "problems.csv", "a csv file with columns 'question' and 'answer'")
	timeLimit := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s\n", *fileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse")
	}

	// fmt.Print(lines)
	problems := parseLines(lines)
	// fmt.Print(problems)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

quizLoop:
	for i, p := range problems {
		fmt.Printf("#%d: %s = ", i+1, p.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTIME'S UP!\nSCORE: %d/%d \n", correct, len(problems))
			break quizLoop
		case answer := <-answerChannel:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("SCORE: %d/%d \n", correct, len(problems))
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
