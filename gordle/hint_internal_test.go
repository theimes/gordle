package gordle

import "testing"

func TestStringConcat(t *testing.T) {
	var fb feedback = make([]hint, 0, 3)
	fb = append(fb, absentCharacter, correctPosition, wrongPosition)

	t.Logf("%s", fb.StringConcat())
}
