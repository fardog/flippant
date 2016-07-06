package mwords

import (
	"bufio"
	"bytes"
	"fmt"
	"sync"

	"github.com/fardog/flippant"
	"github.com/fardog/flippant/mwords/assets"
)

const assetName = "resources/113809of.fic"

var words []string
var once sync.Once

func populateWordList() {
	rw, err := assets.Asset(assetName)

	if err != nil {
		panic(fmt.Sprintf("unable to load asset: %s", assetName))
	}

	scanner := bufio.NewScanner(bytes.NewReader(rw))
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
}

// NewGenerator creates a new flippant.Generator with a provided word list
func NewGenerator() *flippant.Generator {
	once.Do(populateWordList)

	ww := make([]string, len(words))
	copy(ww, words)

	return flippant.NewGenerator(words)
}
