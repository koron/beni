package lexer

import (
	"testing"

	c "github.com/koron/beni/token"
)

func TestGo(t *testing.T) {
	parseCheck(t, Go, `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World\n")
}
`, []interface{}{
		c.Other, "\n",
		c.Keyword, "package", c.Other, " ", c.Name, "main",
		c.Other, "\n\n",

		c.Keyword, "import", c.Other, " ", c.Punctuation, "(",
		c.Other, "\n\t", c.LiteralString, `"fmt"`, c.Other, "\n",
		c.Punctuation, ")",
		c.Other, "\n\n",

		c.Keyword, "func", c.Other, " ", c.Name, "main", c.Punctuation, "(", c.Punctuation, ")", c.Other, " ", c.Punctuation, "{", c.Other, "\n\t",
		c.Name, "fmt", c.Operator, ".", c.Name, "Println", c.Punctuation, "(", c.LiteralString, `"Hello World\n"`, c.Punctuation, ")", c.Other, "\n",
		c.Punctuation, "}", c.Other, "\n",
	}, false)
}
