package engine

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const maxGuess uint8 = 6

type Game struct {
	// configure Game parameters
	//maximum allowed  # of guesses
	ValidWords []string
	// specifics of this Game
	Target    string
	GuessList [maxGuess]string
	Correcter map[string][5][2]bool
	GuessNum  uint8
	Success   bool
}

// update Game type with a new guess

func (e *Game) Guess(guess string) {
	// takes a VALID guess and updates the Game struct with the guess, returns true if the
	// guess matches the Target.
	var err error
	guess, err = e.ConditionGuess(guess)
	if err != nil {
	}
	e.Correcter[guess] = e.GuessCorrecter(guess)
	e.GuessList[e.GuessNum] = guess
	e.GuessNum++
	if guess == e.Target {
		e.Success = true
	}
}

func (e Game) checkWord(guess string) (bool, error) {
	// checks if the word is in the dictionary of valid words
	var validWord bool
	guess = strings.TrimSpace(strings.ToLower(guess))
	var err error
	if utf8.RuneCountInString(guess) != len(guess) {
		err = errors.New("string includes non-ascii charachters. valid guesses must be english letters")
	} else if utf8.RuneCountInString(guess) != 5 {
		err = errors.New("a valid guess must have 5 letters")
	} else {
		for _, wrd := range e.ValidWords {
			if strings.Compare(wrd, guess) == 0 {
				validWord = true
				break
			} else {
				err = errors.New("word not in list")
			}
		}
	}
	return validWord, err
}

func (e Game) ConditionGuess(guess string) (string, error) {
	var validWord bool
	var err error
	guess = strings.TrimSpace(strings.ToLower(guess))
	validWord, err = e.checkWord(guess)
	if err != nil {
		return guess, err
	} else if validWord {
		return guess, nil
	} else {
		return guess, err
	}

}

func WordListReader(fileName string) ([]string, error) {
	fp := filepath.Join("/Users/zach/code/wordgo/assets/data/in", fileName)

	file, err := os.Open(fp)
	if err != nil {

		return make([]string, 1), err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return make([]string, 1), err
	}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func (e Game) GuessCorrecter(guess string) [5][2]bool {
	var checkArr [5][2]bool
	for idx, char1 := range guess {
		if byte(char1) == e.Target[idx] {
			checkArr[idx][0] = true
			checkArr[idx][1] = true
		} else {
			for _, char2 := range e.Target {
				if char1 == char2 {
					checkArr[idx][1] = true
				}
				continue
			}
		}

	}

	return checkArr
}
