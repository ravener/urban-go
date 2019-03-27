# Urban Go
Urban Dictionary API for Go (golang)

[View the GoDoc](https://godoc.org/github.com/pollen5/urban-go)

## Install
```sh
$ go get github.com/pollen5/urban-go
```

## Usage
```go
package main

import (
  "fmt"
  "github.com/pollen5/urban-go"
)

func main() {
  entries, err := urban.Define("term")
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, entry := range entries {
    fmt.Printf("Definition: %s\nExample: %s\nAuthor: %s",
      entry.Definition,
      entry.Example,
      entry.Author,
    )
  }
}
```

### Markdown/HTML
A utility is provided to render a definition into markdown or html
style
```go
urban.Markdown("[some term]")
// => [some term](https://www.urbandictionary.com/define.php?term=some+term)
urban.HTML("[some term]")
// => <a href="https://www.urbandictionary.com/define.php?term=some+term">some term</a>
```

## License
[MIT](LICENSE)
