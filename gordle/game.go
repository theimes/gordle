package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

const wordLength int = 5

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// Game represents a game of gordle
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a playable game
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 😍 You won! You found the solution in %d attempts. The word was: %q\n", currentAttempt, string(g.solution))
			return
		}

		fmt.Printf("%s", computeFeedback(guess, g.solution))

	}

	fmt.Printf("🙁 😜 You've lost. The solution was %q\n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", wordLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s\n", err.Error())
		} else {
			return guess
		}
	}

}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d characters, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacters returns the input string as a slice of uppercased runes
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback verifies every character of the guess against the solution
func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	for i, character := range guess {
		if character == solution[i] {
			result[i] = correctPosition
			used[i] = true
		}
	}

	for i, character := range guess {
		// omit characters that ain't in the solution
		if result[i] != absentCharacter {
			continue
		}
		for j, target := range solution {
			if used[j] {
				continue
			}
			if character == target {
				result[i] = wrongPosition
				used[j] = true
				break
			}
		}
	}
	return result
}
