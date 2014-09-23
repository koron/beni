package lexer

import (
	. "github.com/koron/beni/token"
)

// Go lexer info.
var goInfo = Info{
	Name:        "Go",
	Aliases:     []string{"go", "golang"},
	Filenames:   []string{`.*\.go`},
	Mimetypes:   []string{"text/x-go", "application/x-go", "text/x-gosrc"},
	Description: "The Go programming language (http://golang.org)",
}

var (
	goKeywords = []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}

	goOperators = []string{
		"+=", "++", "+", "&^=", "&^", "&=", "&&", "&", "==", "=",
		"!=", "!", "-=", "--", "-", "|=", "||", "|", "<=", "<-",
		"<<=", "<<", "<", "*=", "*", "^=", "^", ">>=", ">>", ">=",
		">", "/=", "/", ":=", "%", "%=", "...", ".", ":",
	}

	goSeparators = []string{
		"(", ")", "[", "]", "{", "}", ",", ";",
	}

	goTypes = []string{
		"bool", "byte", "complex64", "complex128", "error",
		"float32", "float64", "int8", "int16", "int32",
		"int64", "int", "rune", "string", "uint8",
		"uint16", "uint32", "uint64", "uintptr", "uint",
	}

	goConstants = []string{
		"true", "false", "iota", "nil",
	}

	goFunctions = []string{
		"append", "cap", "close", "complex", "copy",
		"delete", "imag", "len", "make", "new",
		"panic", "print", "println", "real", "recover",
	}
)

var goStates = map[RegexpLexerState][]RegexpLexerRule{
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
			Pattern: "^(?:" + regexpQuoteJoin(goKeywords...) + ")\\b",
			Action:  RegexpEmit(Keyword),
		},
		RegexpLexerRule{
			Name:    "predeclared type",
			Pattern: "^(?:" + regexpQuoteJoin(goTypes...) + ")\\b",
			Action:  RegexpEmit(KeywordType),
		},
		RegexpLexerRule{
			Name:    "predeclared function",
			Pattern: "^(?:" + regexpQuoteJoin(goFunctions...) + ")\\b",
			Action:  RegexpEmit(NameBuiltin),
		},
		RegexpLexerRule{
			Name:    "predeclared constant",
			Pattern: "^(?:" + regexpQuoteJoin(goConstants...) + ")\\b",
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
			Action: RegexpEmit(LiteralNumber),
		},

		// Float literals
		RegexpLexerRule{
			Name: "float lit",
			Pattern: "^(?:" + regexpJoin(
				`\d+(?:[eE][+-]?\d+)`,
				`.\d+(?:[eE][+-]?\d+)?`,
				`\d+\.\d+(?:[eE][+-]?\d+)?`,
			) + ")",
			Action: RegexpEmit(LiteralNumber),
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
			Pattern: "^(?:" + regexpQuoteJoin(goOperators...) + ")",
			Action:  RegexpEmit(Operator),
		},
		RegexpLexerRule{
			Name:    "separator",
			Pattern: "^(?:" + regexpQuoteJoin(goSeparators...) + ")",
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
	return goInfo
}

func (f *goFactory) New() (Lexer, error) {
	return NewRegexpLexer(&RegexpLexerDef{
		Info:   goInfo,
		States: goStates,
	})
}

// Go lexer factory.
var Go = &goFactory{}
