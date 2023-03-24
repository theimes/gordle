package gordle

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "ﻣﺮﺣﺒﺎ",
			want:  []rune("ﻣﺮﺣﺒﺎ"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		//"3 characters in japanese": {
		//	input: "こんに\nこんにちは",
		//	want:  []rune("こんに"),
		//},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, err := New(strings.NewReader(tc.input), []string{"hello"}, 5)
			if err != nil {
				t.Errorf("can't load game: %s", err.Error())
			}

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("readRunes() got = %v, want: %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateLength(t *testing.T) {
	tt := map[string]struct {
		input []rune
		want  error
	}{
		"10 characters in english": {
			input: []rune("HELLOWORLD"),
			want:  errInvalidWordLength,
		},
		"5 characters in english": {
			input: []rune("hello"),
			want:  errInvalidWordLength,
		},
		"3 characters in english": {
			input: []rune("foo"),
			want:  errInvalidWordLength,
		},
		"5 characters in japanese": {
			input: []rune("んにちは"),
			want:  errInvalidWordLength,
		},
		"3 characters in japanese": {
			input: []rune("こんに"),
			want:  errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, err := New(nil, []string{"hello"}, 5)
			if err != nil {
				t.Errorf("can't load game: %s", err.Error())
			}

			got := g.validateGuess(tc.input)
			if got != nil {
				if !errors.Is(got, errInvalidWordLength) {
					t.Errorf("expected %T, got %T", errInvalidWordLength, got)
				}
			}
		})
	}
}

func TestSplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"lowercase input": {
			input: "hello",
			want:  []rune("HELLO"),
		},
		"uppercase input": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"mixed case input": {
			input: "HellO",
			want:  []rune("HELLO"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUppercaseCharacters(tc.input)

			if !slices.Equal(got, tc.want) {
				t.Errorf("expected %s, got %s", string(tc.want), string(got))
			}
		})
	}

}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "hello",
			solution:         "hello",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"two characters wrong position": {
			guess:            "llaaa",
			solution:         "hello",
			expectedFeedback: feedback{wrongPosition, wrongPosition, absentCharacter, absentCharacter, absentCharacter},
		},
		"one right, one wrong": {
			guess:            "laala",
			solution:         "hello",
			expectedFeedback: feedback{wrongPosition, absentCharacter, absentCharacter, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got wrong feedback, expected %v, got %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}

}
