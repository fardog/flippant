package mwords

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fardog/flippant"
	"github.com/fardog/flippant/mwords/assets"
)

const assetName = "resources/113809of.fic"

var words []string

func init() {
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
	ww := make([]string, len(words))
	copy(ww, words)

	return flippant.NewGenerator(words)
}
