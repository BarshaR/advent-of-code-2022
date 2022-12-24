package dayTwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFilePath = "./days/dayTwo/input.txt"

func Run() {
	var rounds []round
	rounds, err := parseInputFile(inputFilePath)
	if err != nil {
		log.Fatal("Error to parsing file", err)
	}

	myScore := 0
	// Calculate my score
	for _, v := range rounds {
		myScore = myScore + v.playerTwoScore
	}

	fmt.Print(rounds)
	fmt.Print(myScore)

}

func parseInputFile(path string) ([]round, error) {
	var rounds []round
	readFile, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal("Unable to read file")
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		rounds = append(rounds, parseRound(fileScanner.Text()))
	}

	return rounds, nil
}

func parseRound(line string) round {
	split := strings.Split(line, " ")
	var round round
	if len(split) == 2 {

		round.playerOneChoice = choiceCodes[split[0]]
		round.playerTwoChoice = choiceCodes[split[1]]

		// calculate winner
		round.winner = calculateWinner(round.playerOneChoice, round.playerTwoChoice)

		// calculate scores
		if round.winner == "draw" {
			round.playerOneScore = 3 + winScore[round.playerOneChoice]
			round.playerTwoScore = 3 + winScore[round.playerTwoChoice]
		} else if round.winner == "playerOne" {
			round.playerOneScore = round.playerOneScore + 6 + winScore[round.playerOneChoice]
			round.playerTwoScore = round.playerTwoScore + winScore[round.playerTwoChoice]
		} else if round.winner == "playerTwo" {
			round.playerTwoScore = round.playerTwoScore + 6 + winScore[round.playerTwoChoice]
			round.playerOneScore = round.playerOneScore + winScore[round.playerOneChoice]
		}

	}
	return round
}

func calculateWinner(playerOneChoice string, playerTwoChoice string) string {
	if playerOneChoice == playerTwoChoice {
		return "draw"
	}

	// Check if playerOneChoice wins against playerTwoChoice
	for _, v := range winsAgainst[playerOneChoice] {
		if v == playerTwoChoice {
			return "playerOne"
		} else {
			return "playerTwo"
		}
	}

	return "Undetermined"
}

var choiceCodes = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var winScore = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var winsAgainst = map[string][]string{
	"Rock":     {"Scissors"},
	"Scissors": {"Paper"},
	"Paper":    {"Rock"},
}

type round struct {
	playerOneChoice string
	playerTwoChoice string
	winner          string // "playerOne", "playerTwo", "draw"
	playerOneScore  int    // I am "playerTwo", score of "playerTwo"
	playerTwoScore  int
}
