package lexer

type RegexpLexerState int32

const (
	Root RegexpLexerState = iota

	JavaClass
	JavaImport
)
