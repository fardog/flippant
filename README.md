# flippant

Go library for retrieving random words from a provided word list.

## Install

```
$ go get github.com/fardog/flippant
```

## Example

```go 
package mypkg

import (
    "fmt"
    "github.com/fardog/flippant"
)

func main() {
    words := []string{"some", "words", "go", "in", "here", "yah"}
    g := flippant.NewGenerator(words)
    
    dest := make([]string, 2)
    g.Words(dest)
    
    fmt.Println(dest) // ["go", "here"]
    
    dest = make([]string, 3)
    g.BoundedUniqueWords(dest, 2, 3)
    
    fmt.Println(dest) // ["in", "yah", "go"]
}
```

## Default Wordlist

There is no default wordlist in the standard `flippant` package, but you can 
use the [`mwords`][] package, which is just `flippant` with a high quality
default wordlist. See [that package][mwords] for further details.

`mwords` also provides an [optional CLI][cli], if you're just looking to get
some random words.

## Other Languages

`flippant` follows a tradition of creating random-word libraries to try out
new languages. You can find similar libraries for the following:

- Node.js: [xkcd-password][]
- Clojure/Clojurescript: [hazard][]

## License

[MIT](./LICENSE)

[xkcd-password]: https://github.com/fardog/node-xkcd-password
[hazard]: https://github.com/fardog/hazard
[mwords]: https://github.com/fardog/flippant/blob/master/mwords
[cli]: https://github.com/fardog/flippant/blob/master/mwords#optional-cli
