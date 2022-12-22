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
	fmt.Println("Hello its me - day 1 :)")

	readFile, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal("Unable to read file")
		return
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	currMax := 0
	currTotal := 0
	var totals []int

	for fileScanner.Scan() {
		if line := fileScanner.Text(); line == "" || line == "\n" {
			// an empty string is found, we have totalled a group.
			// add each group total to an array and sort to find the top 3
			totals = append(totals, currTotal)
			if currTotal > currMax {
				currMax = currTotal
			}
			currTotal = 0
		} else {
			if val, err := strconv.Atoi(line); err == nil {
				currTotal = currTotal + val
			}

		}
		if currTotal > currMax {
			currMax = currTotal
		}
	}

	sort.Ints(totals)
	if len(totals) > 3 {
		fmt.Printf("Top 3 totals - %v, %v, %v", totals[len(totals)-1], totals[len(totals)-2], totals[len(totals)-3])
		fmt.Printf("\nGrant Total of Top 3 - %v\n", totals[len(totals)-1]+totals[len(totals)-2]+totals[len(totals)-3])
	}

	// fmt.Println(currMax)
}
