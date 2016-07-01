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

	into := make([]string, 3)
	l, err := g.Words(into)

	if err != nil {
		t.Error("Received error")
	}
	if l != 3 {
		t.Error("Received unexpected length as return value")
	}

	for _, item := range into {
		if _, found := indexOf(item, words); !found {
			t.Error("Received unexpected item")
		}
	}
}

func TestWord(t *testing.T) {
	g := NewGenerator(words)

	word := g.Word()

	if _, found := indexOf(word, words); !found {
		t.Error("Received unexpected item")
	}
}
