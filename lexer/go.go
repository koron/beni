package lexer

import (
	. "github.com/koron/beni/token"
	"strings"
)

var GO_INFO = Info{
	Name:        "go",
	Aliases:     []string{"go", "golang"},
	Filenames:   []string{"*.go"},
	Mimetypes:   []string{"text/x-go", "application/x-go"},
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

	// Characters
	GO_WHITE_SPACE    = "[\\s\\t\\r\\n]+"
	GO_NEWLINE        = "\\n"
	GO_UNICODE_CHAR   = "[^\\n]"
	GO_UNICODE_LETTER = "[[:alpha:]]"
	GO_UNICODE_DIGIT  = "[[:digit:]]"

	// Letters and digits
	GO_LETTER        = "[[:alpha:]|]"
	GO_DECIMAL_DIGIT = "[0-9]"
	GO_OCTAL_DIGIT   = "[0-7]"
	GO_HEX_DIGIT     = "[0-9A-Fa-f]"

	// Comments
	GO_LINE_COMMENT    = "//[^" + GO_NEWLINE + "]*"
	GO_GENERAL_COMMENT = "(?s:/\\*.*?\\*/)"
	GO_COMMENT         = GO_LINE_COMMENT + "|" + GO_GENERAL_COMMENT

	// Keywords
	GO_KEYWORD = "\\b(?:" + regexpLexerJoin(GO_KEYWORDS) + ")\\b"

	// Identifiers
	GO_IDENTIFIER = "[[:alpha:]]|[[:alpha:][:digit:]]*"

	// Operators and delimiters
	GO_OPERATOR  = regexpLexerJoin(GO_OPERATORS)
	GO_SEPARATOR = regexpLexerJoin(GO_SEPARATORS)

	// Integer literals
	GO_DECIMAL_LIT = "[1-9]" + GO_DECIMAL_DIGIT + "*"
	GO_OCTAL_LIT   = "0" + GO_OCTAL_DIGIT + "*"
	GO_HEX_LIT     = "0[xX]" + GO_HEX_DIGIT + "+"
	GO_INT_LIT     = strings.Join([]string{
		GO_HEX_LIT, GO_DECIMAL_LIT, GO_OCTAL_LIT,
	}, "|")

	// Floating-point literals
	GO_DECIMALS  = GO_DECIMAL_DIGIT + "+"
	GO_EXPONENT  = "[eE][+-]?" + GO_DECIMALS
	GO_FLOAT_LIT = "" +
		GO_DECIMALS + "\\\\." + GO_DECIMALS + "?" + GO_EXPONENT + "?" +
		"|" + GO_DECIMALS + GO_EXPONENT +
		"|" + "\\\\." + GO_DECIMALS + GO_EXPONENT + "?"

	// Imaginary literals
	GO_IMAGINARY_LIT = "(?:" + GO_DECIMALS + "|" + GO_FLOAT_LIT + ")i"

	// Rune literals
	// String literals

	// Predeclared identifiers
	GO_PREDECLARED_TYPES     = "\\b(?:" + regexpLexerJoin(GO_TYPES) + ")\\b"
	GO_PREDECLARED_CONSTANTS = "\\b(?:" + regexpLexerJoin(GO_CONSTANTS) + ")\\b"
	GO_PREDECLARED_FUNCTIONS = "\\b(?:" + regexpLexerJoin(GO_FUNCTIONS) + ")\\b"
)

var GO_STATES = map[RegexpLexerState][]RegexpLexerRule{
	Root: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern: "^(?:" + GO_COMMENT + ")",
			Action:  RegexpEmit(Comment),
		},
		RegexpLexerRule{
			Pattern: "^" + GO_KEYWORD,
			Action:  RegexpEmit(Keyword),
		},
		RegexpLexerRule{
			Pattern: "^" + GO_PREDECLARED_TYPES,
			Action:  RegexpEmit(KeywordType),
		},
		RegexpLexerRule{
			Pattern: "^" + GO_PREDECLARED_FUNCTIONS,
			Action:  RegexpEmit(NameBuiltin),
		},
		RegexpLexerRule{
			Pattern: "^" + GO_PREDECLARED_CONSTANTS,
			Action:  RegexpEmit(NameConstant),
		},
		RegexpLexerRule{
			Pattern: "^" + GO_IMAGINARY_LIT,
			Action:  RegexpEmit(LiteralNumber),
		},
		RegexpLexerRule{
			Pattern: "^(?:" + GO_FLOAT_LIT + ")",
			Action:  RegexpEmit(LiteralNumber),
		},
		RegexpLexerRule{
			Pattern: "^(?:" + GO_INT_LIT + ")",
			Action:  RegexpEmit(LiteralNumber),
		},
		// TODO: add rules
		RegexpLexerRule{
			Pattern: GO_WHITE_SPACE,
			Action:  RegexpEmit(Other),
		},
	},

	GoRawString: []RegexpLexerRule{
		// TODO: support escape sequence
		RegexpLexerRule{
			Pattern: "^[^\"]+",
			Action:  RegexpEmit(LiteralString),
		},
		RegexpLexerRule{
			Pattern: "^\"",
			Action:  RegexpEmitPop(LiteralString),
		},
	},

	GoInterpretedString: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern: "^(?m:^[^`]+)",
			Action:  RegexpEmit(LiteralString),
		},
		RegexpLexerRule{
			Pattern: "^`",
			Action:  RegexpEmitPop(LiteralString),
		},
	},
}
