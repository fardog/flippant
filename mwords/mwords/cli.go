package main

import (
	"fmt"
	"strings"

	"github.com/fardog/flippant/mwords"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Number  uint8 `cli:"n,num" usage:"number of words" dft:"4"`
	Minimum uint8 `cli:"m,min" usage:"minimum word length" dft:"5"`
	Maximum uint8 `cli:"x,max" usage:"maximum word length" dft:"7"`
}

func main() {
	cli.Run(&argT{}, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		g := mwords.NewGenerator()
		dest := make([]string, argv.Number)

		_, err := g.BoundedUniqueWords(dest, int(argv.Minimum), int(argv.Maximum))

		if err != nil {
			return err
		}

		out := strings.Join(dest, " ")
		ctx.String(fmt.Sprintf("%s\n", out))

		return nil
	})
}
