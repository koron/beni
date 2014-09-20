package lexer

import (
	"fmt"
	. "github.com/koron/beni/token"
	"strings"
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
			Behavior: func(c RegexpLexerContext, groups []string) error {
				if err := c.ParseStr(groups[1]); err != nil {
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
		RegexpLexerRule{Pattern: "\\s+", Behavior: RegexpEmit(Text)},
		RegexpLexerRule{
			Pattern:  "//.*?$",
			Behavior: RegexpEmit(CommentSingle),
		},
		RegexpLexerRule{
			Pattern:  "(?s:/\\*.*?\\*/)",
			Behavior: RegexpEmit(CommentMultiline),
		},
		RegexpLexerRule{
			Pattern:  "@" + javaId,
			Behavior: RegexpEmit(NameDecorator),
		},
		RegexpLexerRule{
			Pattern:  "(?:" + strings.Join(javaKeywords, "|") + ")\\b",
			Behavior: RegexpEmit(Keyword),
		},
		RegexpLexerRule{
			Pattern:  "(?:" + strings.Join(javaDeclarations, "|") + ")\\b",
			Behavior: RegexpEmit(KeywordDeclaration),
		},
		RegexpLexerRule{
			Pattern:  "(?:" + strings.Join(javaTypes, "|") + ")\\b",
			Behavior: RegexpEmit(KeywordType),
		},
		RegexpLexerRule{
			Pattern:  "package\\b",
			Behavior: RegexpEmit(KeywordNamespace),
		},
		RegexpLexerRule{
			Pattern:  "(?:true|false|null)\\b",
			Behavior: RegexpEmit(KeywordConstant),
		},
		RegexpLexerRule{
			Pattern:  "(?:class|interface)\\b",
			Behavior: RegexpEmitPush(KeywordDeclaration, JavaClass),
		},
		RegexpLexerRule{
			Pattern:  "import\b",
			Behavior: RegexpEmitPush(KeywordNamespace, JavaImport),
		},
		RegexpLexerRule{
			Pattern:  "(\\\\|\\\"|[^\"])*\"",
			Behavior: RegexpEmit(LiteralString),
		},
		RegexpLexerRule{
			Pattern:  "'(?:\\.|[^\\]|\\u[0-9a-fA-F]{4})'",
			Behavior: RegexpEmit(LiteralStringChar),
		},
		RegexpLexerRule{
			Pattern: "(\\.)(" + javaId + ")",
			Behavior: func(c RegexpLexerContext, groups []string) error {
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
			Pattern:  javaId + ":",
			Behavior: RegexpEmit(NameLabel),
		},
		RegexpLexerRule{
			Pattern:  "\\$?" + javaId,
			Behavior: RegexpEmit(Name),
		},
		RegexpLexerRule{
			Pattern:  "[~^*~%&\\[\\](){}<>\\|+=:;,./?-]",
			Behavior: RegexpEmit(Operator),
		},
		RegexpLexerRule{
			Pattern:  "[0-9][0-9]*\\.[0-9]+([eE][0-9]+)?[fd]?",
			Behavior: RegexpEmit(LiteralNumber),
		},
		RegexpLexerRule{
			Pattern:  "0x[0-9a-fA-F]+",
			Behavior: RegexpEmit(LiteralNumberHex),
		},
		RegexpLexerRule{
			Pattern:  "[0-9]+L?",
			Behavior: RegexpEmit(LiteralNumberInteger),
		},
	},

	JavaClass: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern:  javaSpaces,
			Behavior: RegexpEmit(Text),
		},
		RegexpLexerRule{
			Pattern:  javaId,
			Behavior: RegexpEmitPop(NameClass),
		},
	},

	JavaImport: []RegexpLexerRule{
		RegexpLexerRule{
			Pattern:  javaSpaces,
			Behavior: RegexpEmit(Text),
		},
		RegexpLexerRule{
			Pattern:  "(?i:[a-zA-Z0-9_.]+\\*?)",
			Behavior: RegexpEmitPop(NameNamespace),
		},
	},
}

type javaFactory struct {
}

func (f *javaFactory) Info() Info {
	return javaInfo
}

func (f *javaFactory) New() (Lexer, error) {
	return NewRegexpLexer(&RegexpLexerData{
		Info: javaInfo,
		States: javaStates,
	})
}

var Java = &javaFactory{}
