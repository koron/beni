package lexer

import (
	"testing"

	c "github.com/koron/beni/token"
)

func TestGo(t *testing.T) {
	parseCheck(t, Go, `
package main
`, []interface{}{
		c.Other, "\n",
		c.Keyword, "package", c.Other, " ", c.Name, "main",
		c.Other, "\n",
	}, false)
}
