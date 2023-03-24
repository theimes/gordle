package gordle

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "XAIPE"}
	word := pickWords(corpus)

	if !inCorpus(corpus, string(word)) {
		t.Errorf("expected %q in corpus %v", word, corpus)
	}

}
