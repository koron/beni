package beni

import (
	"strings"
)

var keywords = []string{
	"assert", "break", "case", "catch", "continue", "default", "do", "else",
	"finally", "for", "if", "goto", "instanceof", "new", "return", "switch",
	"this", "throw", "try", "while",
}

var declarations = []string{
	"abstract", "const", "enum", "extends", "final", "implements", "native",
	"private", "protected", "public", "static", "strictfp", "super",
	"synchronized", "throws", "transient", "volatile",
}

var types = []string{
	"boolean", "byte", "char", "double", "float", "int", "long", "short",
	"void",
}

var spaces = "(?m:\\s+)"
var id = "[a-zA-Z_][a-zA-Z0-9_]*"

var JavaLexer = &RegexpLexer{
	Info: LexerInfo{
		Name:        "Java",
		Aliases:     []string{"java"},
		Filenames:   []string{"*.java"},
		Mimetypes:   []string{"text/x-java"},
		Description: "The Java programming language (java.com)",
	},
	Tokens: map[State][]*Rule{
		Root: []*Rule{
			&Rule{Pattern: "\\s+", Token: Text},
			&Rule{Pattern: "//.*?$", Token: CommentSingle},
			&Rule{Pattern: "(?m:/\\*.*?\\*/)", Token: CommentMultiline},
			&Rule{Pattern: "@" + id, Token: NameDecorator},
			&Rule{
				Pattern: "(?:" + strings.Join(keywords, "|") + ")\\b",
				Token:   Keyword,
			},
			&Rule{
				Pattern: "(?:" + strings.Join(declarations, "|") + ")\\b",
				Token:   KeywordDeclaration,
			},
			&Rule{
				Pattern: "(?:" + strings.Join(types, "|") + ")\\b",
				Token:   KeywordType,
			},
			&Rule{Pattern: "package\\b", Token: KeywordNamespace},
			&Rule{
				Pattern: "(?:true|false|null)\\b",
				Token:   KeywordConstant,
			},
			&Rule{
				Pattern: "(?:class|interface)\\b",
				Token:   KeywordDeclaration,
				Next:    Class,
			},
		},
		Class: []*Rule{
			&Rule{
				Pattern: spaces,
				Token:   Text,
			},
			&Rule{
				Pattern: id,
				Token:   NameClass,
				Next:    Pop,
			},
		},
		Import: []*Rule{
			&Rule{
				Pattern: spaces,
				Token:   Text,
			},
			&Rule{
				Pattern: "(?i:[a-zA-Z0-9_.]+\\*?)",
				Token:   NameNamespace,
				Next:    Pop,
			},
		},
	},
}
