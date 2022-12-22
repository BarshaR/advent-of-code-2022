package dayOne

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const inputFilePath = "./days/dayOne/input.txt"

func Run() {
	readFile, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal("Unable to read file")
		return
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	currentTotal := 0
	var totals []int

	for fileScanner.Scan() {
		if line := fileScanner.Text(); line == "" || line == "\n" {
			totals = append(totals, currentTotal)
			currentTotal = 0
		} else {
			if val, err := strconv.Atoi(line); err == nil {
				currentTotal = currentTotal + val
			}

		}
	}

	sort.Ints(totals)
	if len(totals) > 3 {
		fmt.Printf("Top 3 totals - %v, %v, %v", totals[len(totals)-1], totals[len(totals)-2], totals[len(totals)-3])
		fmt.Printf("\nGrand Total of Top 3 - %v\n", totals[len(totals)-1]+totals[len(totals)-2]+totals[len(totals)-3])
	}

}
