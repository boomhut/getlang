# getlang

[![GoDoc](https://godoc.org/github.com/boomhut/getlang?status.svg)](https://godoc.org/github.com/boomhut/getlang) [![Go Report Card](https://goreportcard.com/badge/github.com/boomhut/getlang)](https://goreportcard.com/report/github.com/boomhut/getlang) 

getlang provides fast natural language detection in Go.

## Features

* Offline -- no internet connection required
* Supports [45 languages](https://github.com/boomhut/getlang/blob/master/LANGUAGES.md)
* Provides ISO 639 language codes
* Fast

## Getting started

Installation:
```sh
    go get -u github.com/boomhut/getlang
```

example:
```go
package main

import (
	"fmt"
	"github.com/boomhut/getlang"
)

func main(){
  info := getlang.FromString("Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach")
  fmt.Println(info.LanguageCode(), info.Confidence())
}
```

## Documentation
[getlang on godoc](https://godoc.org/github.com/boomhut/getlang)

## License
[MIT](https://github.com/boomhut/getlang/blob/master/LICENSE)

## Acknowledgements and Citations
* Thanks to [rylans](https://github.com/rylans/getlang) for the original implementation of getlang in Go
* Thanks to [abadojack](https://github.com/abadojack) for the trigram generation logic in whatlanggo
* Cavnar, William B., and John M. Trenkle. "N-gram-based text categorization." Ann arbor mi 48113.2 (1994): 161-175.
