ConfigParser for Go
===================

Implementation of a parser of Python's [ConfigParser] file format for Go.

The format is similar to win INI's, but key/value pairs can be separated by
either `'` or `:`.

Whitespace is trimmed, but newlines are preserved.

[ConfigParser]: http://docs.python.org/library/configparser.html

Example usage:

```go
package main

import (
	"fmt"
	"github.com/mikelikespie/configparser-go/configparser"
)

func main() {
	// configparser.Parse and configparser.ParseFile also work
	config, _ := configparser.ParseString(`[Hello]
foo = bar
boo : baz`)

	foo := config["Hello"]["foo"]

	fmt.Println(foo)
}
```
