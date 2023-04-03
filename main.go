package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main(){
	fileName := flag.String("csv", "problems.csv", "a csv file with columns 'question' and 'answer'" )
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

	correct:=0
	for i, p := range problems {
		fmt.Printf("#%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("SCORE: %d/%d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem{
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

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
