package flippant

import (
	"math/rand"
	"time"
)

// NewGenerator creates a new word generator function
func NewGenerator(words []string) Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{words, r}
}

// Generator is a flippant word generator
type Generator struct {
	Source []string
	Rand   *rand.Rand
}

// Words creates a set of words from the source list into the destination list
func (g *Generator) Words(dest []string) (int, error) {
	ln := len(dest)

	for i := 0; i < ln; i++ {
		dest[i] = g.Word()
	}

	return ln, nil
}

// Word gets a single random word from the source list
func (g Generator) Word() string {
	return g.Source[g.Rand.Intn(len(g.Source))]
}
