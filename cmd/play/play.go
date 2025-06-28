package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zornsLemons/wordGo/pkg/engine"
)

func formatOutputString(cmap []string, guess string) string {
	var s string
	for i := range 5 {
		s = s + cmap[i] + string(guess[i])
	}
	return s
}
func stringColor(checkArr [5][2]bool) []string {
	const colorGreen string = "\033[32m"
	const colorYellow string = "\033[33m"
	const colorReset string = "\033[0m"
	correctionArr := make([]string, 5)
	for i := range 5 {
		if checkArr[i][0] && checkArr[i][1] {
			correctionArr[i] = colorGreen
		} else if checkArr[i][0] || checkArr[i][1] {
			correctionArr[i] = colorYellow
		} else {
			correctionArr[i] = colorReset
		}
	}
	return correctionArr
}

func formatOutputString(cmap []string, guess string) string {
	var s string
	for i := range 5 {
		s = s + cmap[i] + string(guess[i])
	}
	return s
}
func stringColor(checkArr [5][2]bool) []string {
	const colorGreen string = "\033[32m"
	const colorYellow string = "\033[33m"
	const colorReset string = "\033[0m"
	correctionArr := make([]string, 5)
	for i := range 5 {
		if checkArr[i][0] && checkArr[i][1] {
			correctionArr[i] = colorGreen
		} else if checkArr[i][0] || checkArr[i][1] {
			correctionArr[i] = colorYellow
		} else {
			correctionArr[i] = colorReset
		}
	}
	return correctionArr
}
func main() {

	var wordList []string
	var err error
	wordList, err = engine.WordListReader("validWords.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// rn := rand.Intn(len(wordList))
	rn := 14433
	// todo add user input from terminal``

	game := engine.Game{ValidWords: wordList, Target: wordList[rn], GuessNum: 0, Correcter: make(map[string][5][2]bool)}
	fmt.Println(game.Target[4])
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Let's play wordle.")
	fmt.Print("Enter your guess > ")
	guess, _ := reader.ReadString('\n')
	// guess, err = game.ConditionGuess(guess)

	// fmt.Println(err)

	game.Guess(guess)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		fmt.Printf("Guess %v was:", game.GuessNum)
		// fmt.Println(formatOutputString(stringColor(game.Correcter[guess]), guess))
		fmt.Println(game.Correcter[guess])
		fmt.Printf("\n The target word was %v.", game.Target)
		// todo add discord bot
	}

}
