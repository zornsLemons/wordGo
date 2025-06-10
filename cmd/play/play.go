package main

import (
	"fmt"
	"math/rand"

	"github.com/zornsLemons/wordGo/pkg/engine"
)

func main() {

	var wordList []string
	var err error
	wordList, err = engine.WordListReader("validWords.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	rn := rand.Intn(len(wordList))
	// todo add user input from terminal

	game := engine.Game{ValidWords: wordList, Target: wordList[rn], GuessNum: 0, Correcter: make(map[string][]string)}
	fmt.Printf("target: %v, guessNum: %v, guessList: %v, success: %v\n", game.Target, game.GuessNum, game.GuessList, game.Success)
	var guess string = "teach"
	guess, _ = game.ConditionGuess(guess)
	game.Guess(guess)
	fmt.Println(formatOutputString(game.GuessCorrecter(guess), guess))
	guess = "Crane"
	guess, _ = game.ConditionGuess(guess)
	game.Guess(guess)
	fmt.Println(formatOutputString(game.GuessCorrecter(guess), guess))
	fmt.Printf("\n The target word was %v.", game.Target)
	// todo add discord bot
}
func formatOutputString(cmap []string, guess string) string {
	var s string
	for i := range 5 {
		s = s + cmap[i] + string(guess[i])
	}
	return s
}
