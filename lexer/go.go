package lexer

import (
	. "github.com/koron/beni/token"
)

var GO_INFO = Info{
	Name:        "go",
	Aliases:     []string{"go", "golang"},
	Filenames:   []string{"*.go"},
	Mimetypes:   []string{"text/x-go", "application/x-go", "text/x-gosrc"},
	Description: "The Go programming language (http://golang.org)",
}

var (
	GO_KEYWORDS = []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}

	GO_OPERATORS = []string{
		"+=", "++", "+", "&^=", "&^", "&=", "&&", "&", "==", "=",
		"!=", "!", "-=", "--", "-", "|=", "||", "|", "<=", "<-",
		"<<=", "<<", "<", "*=", "*", "^=", "^", ">>=", ">>", ">=",
		">", "/=", "/", ":=", "%", "%=", "...", ".", ":",
	}

	GO_SEPARATORS = []string{
		"(", ")", "[", "]", "{", "}", ",", ";",
	}

	GO_TYPES = []string{
		"bool", "byte", "complex64", "complex128", "error",
		"float32", "float64", "int8", "int16", "int32",
		"int64", "int", "rune", "string", "uint8",
		"uint16", "uint32", "uint64", "uintptr", "uint",
	}

	GO_CONSTANTS = []string{
		"true", "false", "iota", "nil",
	}

	GO_FUNCTIONS = []string{
		"append", "cap", "close", "complex", "copy",
		"delete", "imag", "len", "make", "new",
		"panic", "print", "println", "real", "recover",
	}
)

var GO_STATES = map[RegexpLexerState][]RegexpLexerRule{
	Root: []RegexpLexerRule{
		// Comments
		RegexpLexerRule{
			Name:    "line comment",
			Pattern: "^//[^\\n]*",
			Action:  RegexpEmit(Comment),
		},
		RegexpLexerRule{
			Name:    "general comment",
			Pattern: "^(?s:/\\*.*?\\*/)",
			Action:  RegexpEmit(Comment),
		},

		// Keywords
		RegexpLexerRule{
			Name:    "keyword",
			Pattern: "^(?:" + regexpQuoteJoin(GO_KEYWORDS...) + ")\\b",
			Action:  RegexpEmit(Keyword),
		},
		RegexpLexerRule{
			Name:    "predeclared type",
			Pattern: "^(?:" + regexpQuoteJoin(GO_TYPES...) + ")\\b",
			Action:  RegexpEmit(KeywordType),
		},
		RegexpLexerRule{
			Name:    "predeclared function",
			Pattern: "^(?:" + regexpQuoteJoin(GO_FUNCTIONS...) + ")\\b",
			Action:  RegexpEmit(NameBuiltin),
		},
		RegexpLexerRule{
			Name:    "predeclared constant",
			Pattern: "^(?:" + regexpQuoteJoin(GO_CONSTANTS...) + ")\\b",
			Action:  RegexpEmit(NameConstant),
		},

		// Literals (except strings)
		RegexpLexerRule{
			Name: "imaginary lit",
			Pattern: "^(?:" + regexpJoin(
				"\\d+i",
				`\d+(?:[eE][+-]?\d+)i`,
				`.\d+(?:[eE][+-]?\d+)?i`,
				`\d+\.\d+(?:[eE][+-]?\d+)?i`,
			) + ")",
			Action:  RegexpEmit(LiteralNumber),
		},

		// Float literals
		RegexpLexerRule{
			Name: "float lit",
			Pattern: "^(?:" + regexpJoin(
				`\d+(?:[eE][+-]?\d+)`,
				`.\d+(?:[eE][+-]?\d+)?`,
				`\d+\.\d+(?:[eE][+-]?\d+)?`,
			) + ")",
			Action:  RegexpEmit(LiteralNumber),
		},

		// Integer literals
		RegexpLexerRule{
			Name:    "octal lit",
			Pattern: "^0[0-7]+",
			Action:  RegexpEmit(LiteralNumberHex),
		},
		RegexpLexerRule{
			Name:    "hex lit",
			Pattern: "^0[xX][[:xdigit:]]+",
			Action:  RegexpEmit(LiteralNumberHex),
		},
		RegexpLexerRule{
			Name:    "decimal lit",
			Pattern: "^(?:0|[1-9]\\d*)",
			Action:  RegexpEmit(LiteralNumberInteger),
		},

		// Character literal
		RegexpLexerRule{
			Name: "char lit",
			Pattern: "^'(?:" + regexpJoin(
				"\\\\[abfnrtv'\"]",
				"\\\\u[[:xdigit:]]{4}",
				"\\\\U[[:xdigit:]]{8}",
				"\\\\x[[:xdigit:]]{2}",
				"\\\\[0-7]{3}",
				"[^\\\\]",
			) + ")'",
			Action: RegexpEmit(LiteralStringChar),
		},

		// Operators and separators
		RegexpLexerRule{
			Name:    "operator",
			Pattern: "^(?:" + regexpQuoteJoin(GO_OPERATORS...) + ")",
			Action:  RegexpEmit(Operator),
		},
		RegexpLexerRule{
			Name:    "separator",
			Pattern: "^(?:" + regexpQuoteJoin(GO_SEPARATORS...) + ")",
			Action:  RegexpEmit(Punctuation),
		},

		// Identifiers
		RegexpLexerRule{
			Name:    "identifier",
			Pattern: "^[\\pL][\\pL\\d_]*",
			Action:  RegexpEmit(Name),
		},

		// Strings
		RegexpLexerRule{
			Name:    "raw string lit",
			Pattern: "^(?s:`[^`]*`)",
			Action:  RegexpEmit(LiteralString),
		},
		RegexpLexerRule{
			Name:    "interpreted string lit",
			Pattern: `^"(?:\\\\|\\"|[^"])*"`,
			Action:  RegexpEmit(LiteralString),
		},

		// Others
		RegexpLexerRule{
			Pattern: "^\\s+",
			Action:  RegexpEmit(Other),
		},
	},
}

type goFactory struct {
}

func (f *goFactory) Info() Info {
	return GO_INFO
}

func (f *goFactory) New() (Lexer, error) {
	return NewRegexpLexer(&RegexpLexerDef{
		Info:   GO_INFO,
		States: GO_STATES,
	})
}

var Go = &goFactory{}
