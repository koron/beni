package lexer

type RegexpLexerState int32

const (
	Root RegexpLexerState = iota

	JavaClass
	JavaImport
)

func (s RegexpLexerState) String() string {
	switch s {
	case Root:
		return "Root"

	case JavaClass:
		return "JavaClass"
	case JavaImport:
		return "JavaImport"

	default:
		return "UNKNOWN"
	}
}
