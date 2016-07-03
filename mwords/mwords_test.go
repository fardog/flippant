package mwords

import (
	"testing"
)

func TestNewGenerator(t *testing.T) {
	_, err := NewGenerator()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestGeneratorParameters(t *testing.T) {
	g, err := NewGenerator()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	words := make([]string, 1000)

	num, err := g.BoundedUniqueWords(words, 3, 7)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if num != 1000 {
		t.Errorf("unexpected return value: %d", num)
	}

	for _, word := range words {
		if l := len(word); l < 3 || l > 7 {
			t.Errorf("unexpected word length: %d for word %s", l, word)
		}
	}
}
