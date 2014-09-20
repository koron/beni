package beni

import (
	"fmt"
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

var spaces = "(?s:\\s+)"
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
			&Rule{
				Pattern: "^" +
					"(\\s*(?:[A-Za-z_][0-9A-Za-z_.\\[\\]]*\\s+)+?)" +
					"([A-Za-z_][0-9A-Za-z_]*)" +
					"(\\s*)(\\()",
				Behavior: func(m StateMachine, groups []string) error {
					// XXX: Specify JavaLexer
					if err := m.Delegate(nil, groups[0]); err != nil {
						return err
					}
					if err := m.Emit(NameFunction, groups[1]); err != nil {
						return err
					}
					if err := m.Emit(Text, groups[2]); err != nil {
						return err
					}
					if err := m.Emit(Punctuation, groups[3]); err != nil {
						return err
					}
					return nil
				},
			},
			&Rule{Pattern: "\\s+", Token: Text},
			&Rule{Pattern: "//.*?$", Token: CommentSingle},
			&Rule{Pattern: "(?s:/\\*.*?\\*/)", Token: CommentMultiline},
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
			&Rule{
				Pattern: "import\b",
				Token:   KeywordNamespace,
				Next:    Import,
			},
			&Rule{
				Pattern: "(\\\\|\\\"|[^\"])*\"",
				Token:   LiteralString,
			},
			&Rule{
				Pattern: "'(?:\\.|[^\\]|\\u[0-9a-fA-F]{4})'",
				Token:   LiteralStringChar,
			},
			&Rule{
				Pattern: "(\\.)(" + id + ")",
				Behavior: func(m StateMachine, groups []string) error {
					if len(groups) != 2 {
						return fmt.Errorf("expected 2 groups, acutual %d",
							len(groups))
					}
					if err := m.Emit(Operator, groups[0]); err != nil {
						return err
					}
					if err := m.Emit(NameAttribute, groups[1]); err != nil {
						return err
					}
					return nil
				},
			},
			&Rule{
				Pattern: id + ":",
				Token:   NameLabel,
			},
			&Rule{
				Pattern: "\\$?" + id,
				Token:   Name,
			},
			&Rule{
				Pattern: "[~^*~%&\\[\\](){}<>\\|+=:;,./?-]",
				Token:   Operator,
			},
			&Rule{
				Pattern: "[0-9][0-9]*\\.[0-9]+([eE][0-9]+)?[fd]?",
				Token:   LiteralNumber,
			},
			&Rule{
				Pattern: "0x[0-9a-fA-F]+",
				Token:   LiteralNumberHex,
			},
			&Rule{
				Pattern: "[0-9]+L?",
				Token:   LiteralNumberInteger,
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
