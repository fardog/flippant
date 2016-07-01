package flippant

import (
	"math"
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
	for i := 0; i < len(dest); i++ {
		dest[i] = g.getRandomWord()
	}

	return len(dest), nil
}

func (g Generator) getRandomWord() string {
	return g.Source[g.boundedInt(len(g.Source))]
}

func (g Generator) boundedInt(max int) uint {
	n := g.Rand.Float64()

	return uint(math.Floor(n * float64(max)))
}
