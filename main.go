package main

import (
	"fmt"
	"os"

	"github.com/theimes/gordle/gordle"
)

func main() {
	corpus, err := gordle.ReadCorpus("corpus/corpus.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err.Error())
		os.Exit(1)
	}

	g, err := gordle.New(os.Stdin, corpus, 5)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "can't load gordle: %s", err.Error())
		os.Exit(1)
	}

	g.Play()
}
