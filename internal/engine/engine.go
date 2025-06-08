package engine

type Game struct {
	target    string
	guessList []string
	guessNum  uint8
	params    GameParams
}

type GameParams struct {
	maxGuess   uint8 //maximum allowed of guesses
	validWords []string
}

// update Game type with a new guess
func initGameParams(g GameParams) {
	g.maxGuess = 6

	// g.validWords = make([]string, )
}
func initGame() {

}

func Guess(game Game, guess string) {
	guess.guessList = append(game.guessList, guess)
	game.guessNum++

}
