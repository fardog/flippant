# mwords

Go library for retrieving random words from a high-quality wordlist. The
`mwords` package is [`flippant`][flippant] with a default wordlist included.

## Install

```
$ go get github.com/fardog/flippant/mwords
```

## Example

```go
package mypkg

import (
    "fmt"
    "github.com/fardog/flippant/mwords"
)

func main() {
    g := mwords.NewGenerator()
    
    dest := make([]string, 4)
    g.BoundedUniqueWords(dest, 5, 7)
    
    fmt.Println(dest) // ["vastly", "veiled", "cobby", "myopia"]
}
```

The wordlist is packed into the application code using [`go-bindata`][bindata].
If you require control on how and when words are loaded, you are encouraged to
build your own implementation using [`flippant`][flippant] directly.

## Optional CLI

`mwords` has an optional CLI which exposes the default wordlist. It can be a
useful tool if you're just looking to get some random words. To install:

```
$ go get github.com/fardog/flippant/mwords/...
```

You can then run `mwords --help` from your shell to see all available options.

## License

### Code

[MIT](./LICENSE)

### Word List

[Public Domain](./assets/GW_LICENSE)

[flippant]: https://github.com/fardog/flippant
[bindata]: https://github.com/jteeuwen/go-bindata
