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

## API

Soon.

## Other Languages

`flippant` follows a tradition of creating random-word libraries to try out
new languages. You can find similar libraries for the following:

- Node.js: [xkcd-password][]
- Clojure/Clojurescript: [hazard][]

## License

[MIT](./LICENSE)

[xkcd-password]: https://github.com/fardog/xkcd-password
[hazard]: https://github.com/fardog/hazard
