package lexer

import (
	"fmt"
	"strings"
	. "github.com/koron/beni/token"
)

var javaInfo = Info{
	Name:        "Java",
	Aliases:     []string{"java"},
	Filenames:   []string{"*.java"},
	Mimetypes:   []string{"text/x-java"},
	Description: "The Java programming language (java.com)",
}

var javaKeywords = []string{
	"assert", "break", "case", "catch", "continue", "default", "do", "else",
	"finally", "for", "if", "goto", "instanceof", "new", "return", "switch",
	"this", "throw", "try", "while",
}

var javaDeclarations = []string{
	"abstract", "const", "enum", "extends", "final", "implements", "native",
	"private", "protected", "public", "static", "strictfp", "super",
	"synchronized", "throws", "transient", "volatile",
}

var javaTypes = []string{
	"boolean", "byte", "char", "double", "float", "int", "long", "short",
	"void",
}

var javaSpaces = "(?s:\\s+)"

var javaId = "[a-zA-Z_][a-zA-Z0-9_]*"

var javaStates = map[RegexpLexerState][]RegexpLexerRule{
	Root: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern: "^" +
				"(\\s*(?:[A-Za-z_][0-9A-Za-z_.\\[\\]]*\\s+)+?)" +
				"(" + javaId + ")" +
				"(\\s*)(\\()",
			Action: func(c RegexpLexerContext, groups []string) error {
				if err := c.ParseString(groups[1]); err != nil {
					return err
				}
				if err := c.Emit(NameFunction, groups[2]); err != nil {
					return err
				}
				if err := c.Emit(Text, groups[3]); err != nil {
					return err
				}
				return c.Emit(Punctuation, groups[4])
			},
		},
		RegexpLexerRule{Pattern: "^\\s+", Action: RegexpEmit(Text)},
		RegexpLexerRule{
			Pattern: "^//.*?$",
			Action:  RegexpEmit(CommentSingle),
		},
		RegexpLexerRule{
			Pattern: "^(?s:/\\*.*?\\*/)",
			Action:  RegexpEmit(CommentMultiline),
		},
		RegexpLexerRule{
			Pattern: "^@" + javaId,
			Action:  RegexpEmit(NameDecorator),
		},
		RegexpLexerRule{
			Pattern: "^(?:" + strings.Join(javaKeywords, "|") + ")\\b",
			Action:  RegexpEmit(Keyword),
		},
		RegexpLexerRule{
			Pattern: "^(?:" + strings.Join(javaDeclarations, "|") + ")\\b",
			Action:  RegexpEmit(KeywordDeclaration),
		},
		RegexpLexerRule{
			Pattern: "^(?:" + strings.Join(javaTypes, "|") + ")\\b",
			Action:  RegexpEmit(KeywordType),
		},
		RegexpLexerRule{
			Pattern: "^package\\b",
			Action:  RegexpEmit(KeywordNamespace),
		},
		RegexpLexerRule{
			Pattern: "^(?:true|false|null)\\b",
			Action:  RegexpEmit(KeywordConstant),
		},
		RegexpLexerRule{
			Pattern: "^(?:class|interface)\\b",
			Action:  RegexpEmitPush(KeywordDeclaration, JavaClass),
		},
		RegexpLexerRule{
			Pattern: "^import\b",
			Action:  RegexpEmitPush(KeywordNamespace, JavaImport),
		},
		RegexpLexerRule{
			Pattern: "^\"(\\\\|\\\"|[^\"])*\"",
			Action:  RegexpEmit(LiteralString),
		},
		RegexpLexerRule{
			Pattern: "^'(?:\\.|[^\\]|\\\\u[0-9a-fA-F]{4})'",
			Action:  RegexpEmit(LiteralStringChar),
		},
		RegexpLexerRule{
			Pattern: "^(\\.)(" + javaId + ")",
			Action: func(c RegexpLexerContext, groups []string) error {
				if len(groups) != 3 {
					return fmt.Errorf("expected 3 groups, acutual %d",
						len(groups))
				}
				if err := c.Emit(Operator, groups[1]); err != nil {
					return err
				}
				return c.Emit(NameAttribute, groups[2])
			},
		},
		RegexpLexerRule{
			Pattern: "^" + javaId + ":",
			Action:  RegexpEmit(NameLabel),
		},
		RegexpLexerRule{
			Pattern: "^\\$?" + javaId,
			Action:  RegexpEmit(Name),
		},
		RegexpLexerRule{
			Pattern: "^[~^*~%&\\[\\](){}<>\\|+=:;,./?-]",
			Action:  RegexpEmit(Operator),
		},
		RegexpLexerRule{
			Pattern: "^[0-9][0-9]*\\.[0-9]+([eE][0-9]+)?[fd]?",
			Action:  RegexpEmit(LiteralNumber),
		},
		RegexpLexerRule{
			Pattern: "^0x[0-9a-fA-F]+",
			Action:  RegexpEmit(LiteralNumberHex),
		},
		RegexpLexerRule{
			Pattern: "^[0-9]+L?",
			Action:  RegexpEmit(LiteralNumberInteger),
		},
	},

	JavaClass: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern: "^" + javaSpaces,
			Action:  RegexpEmit(Text),
		},
		RegexpLexerRule{
			Pattern: "^" + javaId,
			Action:  RegexpEmitPop(NameClass),
		},
	},

	JavaImport: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern: "^" + javaSpaces,
			Action:  RegexpEmit(Text),
		},
		RegexpLexerRule{
			Pattern: "^(?i:[a-zA-Z0-9_.]+\\*?)",
			Action:  RegexpEmitPop(NameNamespace),
		},
	},
}

type javaFactory struct {
}

func (f *javaFactory) Info() Info {
	return javaInfo
}

func (f *javaFactory) New() (Lexer, error) {
	return NewRegexpLexer(&RegexpLexerDef{
		Info:   javaInfo,
		States: javaStates,
	})
}

var Java = &javaFactory{}
