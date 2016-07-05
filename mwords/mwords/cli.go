package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fardog/flippant"
	"github.com/fardog/flippant/mwords"
	"github.com/mkideal/cli"
)

var app = &cli.Command{
	Name: os.Args[0],
	Desc: "Random word picker",
	Text: "Get random unique words from a provided word list.",
	Argv: func() interface{} { return new(argT) },
	Fn:   mwordscli,
}

type argT struct {
	cli.Helper
	Number    uint32 `cli:"n,num" usage:"number of words" dft:"4"`
	Minimum   uint8  `cli:"m,min" usage:"minimum word length" dft:"5"`
	Maximum   uint8  `cli:"x,max" usage:"maximum word length" dft:"7"`
	FilePath  string `cli:"f,filepath" usage:"optional path to a newline-delimited word file"`
	Separator string `cli:"s,separator" usage:"separator between words" dft:" "`
}

func readWordFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func mwordscli(ctx *cli.Context) error {
	argv := ctx.Argv().(*argT)

	var g *flippant.Generator

	if len(argv.FilePath) > 0 {
		words, err := readWordFile(argv.FilePath)
		if err != nil {
			return err
		}

		g = flippant.NewGenerator(words)
	} else {
		g = mwords.NewGenerator()
	}
	dest := make([]string, argv.Number)

	_, err := g.BoundedUniqueWords(dest, int(argv.Minimum), int(argv.Maximum))

	if err != nil {
		return err
	}

	out := strings.Join(dest, argv.Separator)
	ctx.String(fmt.Sprintf("%s\n", out))

	return nil
}

func main() {
	cli.SetUsageStyle(cli.ManualStyle)

	if err := app.RunWith(os.Args[1:], os.Stderr, nil); err != nil {
		fmt.Printf("%v\n", err)
	}
}
