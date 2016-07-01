package flippant

import (
	"fmt"
	"math/rand"
	"time"
)

// NewGenerator creates a new word generator function
func NewGenerator(words []string) Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{words, len(words), r}
}

// Generator is a flippant word generator
type Generator struct {
	source []string
	len    int
	r      *rand.Rand
}

// Words creates an array of words from the source list into the
// destination list
func (g Generator) Words(dest []string) (int, error) {
	ln := len(dest)

	for i := 0; i < ln; i++ {
		dest[i] = g.Word()
	}

	return ln, nil
}

// UniqueWords creates a set of unique words from the source list
func (g Generator) UniqueWords(dest []string) (int, error) {
	dl := len(dest)

	if dl > g.len {
		return 0, fmt.Errorf(
			"destination is larger than word source; %d > %d", dl, g.len,
		)
	}

	sel := g.r.Perm(g.len)

	for i, idx := range sel[:dl] {
		dest[i] = g.source[idx]
	}

	return dl, nil
}

// Word gets a single random word from the source list
func (g Generator) Word() string {
	return g.source[g.r.Intn(g.len)]
}
