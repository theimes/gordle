package main

import (
	"os"

	"github.com/theimes/gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin, "hello", 5)

	g.Play()
}
