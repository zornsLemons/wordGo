package main

import (
	"fmt"

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
	// rn := rand.Intn(len(wordList))
	rn := 14433

	game := engine.Game{ValidWords: wordList, Target: wordList[rn], GuessNum: 0, Correcter: make(map[string][5][2]bool)}

	fmt.Println("Hello! Let's play wordle.")

	while := true

	for while {

		while = game.GameLoop()
	}
	fmt.Printf("\n The target word was %v .", game.Target)

	// TODO: Implement discord bot

}
