package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type corpusError string

func (e corpusError) Error() string {
	return string(e)
}

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// split the input on white space
	words := strings.Fields(string(data))

	return words, nil
}

// pickWords returns a random word from the corpus
func pickWords(corpus []string) []rune {

	chaos := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	index := chaos.Intn(len(corpus))

	return []rune(corpus[index])
}
