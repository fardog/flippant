package flippant

import "testing"

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
