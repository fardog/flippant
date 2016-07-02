package flippant

import (
	"math/rand"
	"testing"
)

var words = []string{"a", "b", "c", "d", "e"}

func indexOf(n string, h []string) (int, bool) {
	for i, item := range h {
		if n == item {
			return i, true
		}
	}

	return 0, false
}

func TestNewGenerator(t *testing.T) {
	_ = NewGenerator(words)
}

func TestWords(t *testing.T) {
	g := NewGenerator(words)

	dest := make([]string, 3)
	l, err := g.Words(dest)

	if err != nil {
		t.Errorf("received error: %s", err)
	}
	if l != 3 {
		t.Errorf("unexpected length: %d", l)
	}

	for _, item := range dest {
		if _, found := indexOf(item, words); !found {
			t.Errorf("unexpected item: %s", item)
		}
	}
}

func TestWord(t *testing.T) {
	g := NewGenerator(words)

	word := g.Word()

	if _, found := indexOf(word, words); !found {
		t.Errorf("unexpected item: %s", word)
	}
}

func TestUniqueWords(t *testing.T) {
	g := NewGenerator(words)

	dest := make([]string, 3)
	l, err := g.UniqueWords(dest)

	if err != nil {
		t.Errorf("received error: %s", err)
	}
	if l != 3 {
		t.Errorf("unexpected length: %d", l)
	}

	set := make(map[string]bool)

	for _, w := range dest {
		if _, found := indexOf(w, words); !found {
			t.Errorf("unexpected item: %s", w)
		}
		if _, ok := set[w]; ok {
			t.Errorf("duplicate item: %s", w)
		}
		set[w] = true
	}
}

func TestUniqueWordsError(t *testing.T) {
	g := NewGenerator(words)

	dest := make([]string, len(words)+1)
	_, err := g.UniqueWords(dest)

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewGeneratorWithRand(t *testing.T) {
	r1 := rand.New(rand.NewSource(1))
	r2 := rand.New(rand.NewSource(1))

	g1 := NewGeneratorWithRand(words, r1)
	g2 := NewGeneratorWithRand(words, r2)

	for i := 0; i < 10; i++ {
		if w1, w2 := g1.Word(), g2.Word(); w1 != w2 {
			t.Errorf("expected identical words, got: %s, %s", w1, w2)
		}
	}
}
