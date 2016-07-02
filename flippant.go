package flippant

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// NewGenerator creates a new word generator function
func NewGenerator(words []string) *Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return NewGeneratorWithRand(words, r)
}

// NewGeneratorWithRand creates a new word generator with a pre-seeded random
func NewGeneratorWithRand(words []string, r *rand.Rand) *Generator {
	ww := make([]string, len(words))
	copy(ww, words)
	sort.Sort(ByLength(ww))

	return &Generator{ww, len(ww), r, MakeLengthMap(ww)}
}

// Generator is a flippant word generator
type Generator struct {
	source []string
	len    int
	r      *rand.Rand
	lm     map[int]int
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

// BoundedWords creates an array of words from the source list, where word
// length is bounded between min and max (inclusive)
func (g Generator) BoundedWords(dest []string, min, max int) (int, error) {
	dl := len(dest)

	words, err := g.sliceForBounds(min, max)
	if err != nil {
		return 0, err
	}

	wl := len(words)

	for i := 0; i < dl; i++ {
		dest[i] = words[g.r.Intn(wl)]
	}

	return dl, nil
}

// BoundedUniqueWords creates an array of unique words from the source list,
// where word length is bounded between min and max (inclusive)
func (g Generator) BoundedUniqueWords(dest []string, min, max int) (int, error) {
	dl := len(dest)

	words, err := g.sliceForBounds(min, max)
	if err != nil {
		return 0, err
	}

	wl := len(words)

	if dl > wl {
		return 0, fmt.Errorf(
			"destination is larger than candidate word source; %d > %d", dl, wl,
		)
	}

	sel := g.r.Perm(wl)

	for i, idx := range sel[:dl] {
		dest[i] = words[idx]
	}

	return dl, nil
}

// Word gets a single random word from the source list
func (g Generator) Word() string {
	return g.source[g.r.Intn(g.len)]
}

func (g Generator) sliceForBounds(min, max int) ([]string, error) {
	var keys []int

	if min > max {
		return nil, fmt.Errorf("invalid bounds: min %d > max %d", min, max)
	}

	for k := range g.lm {
		keys = append(keys, k)
	}

	fm := 0
	for _, k := range keys {
		if k >= min {
			fm = g.lm[k]
			break
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	fx := g.len
	for _, k := range keys {
		if k <= max {
			break
		}
		fx = g.lm[k]
	}

	return g.source[fm:fx], nil
}
