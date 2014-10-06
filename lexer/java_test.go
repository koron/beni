package lexer

import (
	"testing"

	c "github.com/koron/beni/token"
)

func TestJava1(t *testing.T) {
	parseCheck(t, Java, `
package net.kaoriya.beni;

/**
 * Lexer class
 */
public class Lexer {
	public Lexer() {}
}
`, []interface{}{
		c.Text, "\n",
		c.KeywordNamespace, "package", c.Text, " ", c.Name, "net",
		c.Operator, ".", c.NameAttribute, "kaoriya", c.Operator, ".",
		c.NameAttribute, "beni", c.Operator, ";",
		c.Text, "\n\n",

		c.CommentMultiline, "/**\n * Lexer class\n */", c.Text, "\n",
		c.KeywordDeclaration, "public", c.Text, " ",
		c.KeywordDeclaration, "class", c.Text, " ",
		c.NameClass, "Lexer", c.Text, " ",
		c.Operator, "{", c.Text, "\n\t",

		c.KeywordDeclaration, "public", c.Text, " ",
		c.NameFunction, "Lexer", c.Text, "",
		c.Punctuation, "(", c.Operator, ")", c.Text, " ",
		c.Operator, "{", c.Operator, "}", c.Text, "\n",

		c.Operator, "}", c.Text, "\n",
	}, false)
}
