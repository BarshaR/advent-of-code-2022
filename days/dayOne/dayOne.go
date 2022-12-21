package dayOne

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for fileScanner.Scan() {
		if line := fileScanner.Text(); line == "" || line == "\n" {
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

	fmt.Println(currMax)
}
