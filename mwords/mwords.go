package mwords

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fardog/flippant"
)

const assetName = "113809of.fic"

// NewGenerator creates a new flippant.Generator with a provided word list
func NewGenerator() (*flippant.Generator, error) {
	rw, err := Asset(assetName)

	if err != nil {
		return nil, fmt.Errorf("unable to load asset: %s", assetName)
	}

	var words []string

	scanner := bufio.NewScanner(bytes.NewReader(rw))
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return flippant.NewGenerator(words), nil
}
