package beni

type Token int32

type TokenData struct {
	Token     Token
	Name      string
	ShortName string
	Parent    Token
}

const (
	Text Token = iota
	TextWhitespace

	Error

	Other

	Keyword
	KeywordConstant
	KeywordDeclaration
	KeywordNamespace
	KeywordPseudo
	KeywordReserved
	KeywordType
	KeywordVariable

	Name
	NameAttribute
	NameBuiltin
	NameBuiltinPseudo
	NameClass
	NameConstant
	NameDecorator
	NameEntity
	NameException
	NameFunction
	NameProperty
	NameLabel
	NameNamespace
	NameOther
	NameTag
	NameVariable
	NameVariableClass
	NameVariableGlobal
	NameVariableInstance

	Literal
	LiteralDate
	LiteralString
	LiteralStringBacktick
	LiteralStringChar
	LiteralStringDoc
	LiteralStringDouble
	LiteralStringEscape
	LiteralStringHeredoc
	LiteralStringInterpol
	LiteralStringOther
	LiteralStringRegex
	LiteralStringSingle
	LiteralStringSymbol
	LiteralNumber
	LiteralNumberFloat
	LiteralNumberHex
	LiteralNumberInteger
	LiteralNumberIntegerLong
	LiteralNumberOct
	LiteralNumberBin
	LiteralNumberOther

	Operator
	OperatorWord

	Punctuation
	PunctuationIndicator

	Comment
	CommentDoc
	CommentMultiline
	CommentPreproc
	CommentSingle
	CommentSpecial

	Generic
	GenericDeleted
	GenericEmph
	GenericError
	GenericHeading
	GenericInserted
	GenericOutput
	GenericPrompt
	GenericStrong
	GenericSubheading
	GenericTraceback
	GenericLineno
)

var Tokens = []TokenData{
	TokenData{Text, "Text", "", 0},
	TokenData{TextWhitespace, "Whitespace", "w", Text},

	TokenData{Error, "Error", "err", 0},
	TokenData{Other, "Other", "x", 0},

	TokenData{Keyword, "Keyword", "k", 0},
	TokenData{KeywordConstant, "Constant", "kc", Keyword},
	TokenData{KeywordDeclaration, "Declaration", "kd", Keyword},
	TokenData{KeywordNamespace, "Namespace", "kn", Keyword},
	TokenData{KeywordPseudo, "Pseudo", "kp", Keyword},
	TokenData{KeywordReserved, "Reserved", "kr", Keyword},
	TokenData{KeywordType, "Type", "kt", Keyword},
	TokenData{KeywordVariable, "Variable", "kv", Keyword},

	TokenData{Name, "Name", "n", 0},
	TokenData{NameAttribute, "Attribute", "na", Name},
	TokenData{NameBuiltin, "Builtin", "nb", Name},
	TokenData{NameBuiltinPseudo, "Pseudo", "bp", NameBuiltin},
	TokenData{NameClass, "Class", "nc", Name},
	TokenData{NameConstant, "Constant", "no", Name},
	TokenData{NameDecorator, "Decorator", "nd", Name},
	TokenData{NameEntity, "Entity", "ni", Name},
	TokenData{NameException, "Exception", "ne", Name},
	TokenData{NameFunction, "Function", "nf", Name},
	TokenData{NameProperty, "Property", "py", Name},
	TokenData{NameLabel, "Label", "nl", Name},
	TokenData{NameNamespace, "Namespace", "nn", Name},
	TokenData{NameOther, "Other", "nx", Name},
	TokenData{NameTag, "Tag", "nt", Name},
	TokenData{NameVariable, "Variable", "nv", Name},
	TokenData{NameVariableClass, "Class", "vc", NameVariable},
	TokenData{NameVariableGlobal, "Global", "vg", NameVariable},
	TokenData{NameVariableInstance, "Instance", "vi", NameVariable},

	TokenData{Literal, "Literal", "l", 0},
	TokenData{LiteralDate, "Date", "ld", Literal},
	TokenData{LiteralString, "String", "s", Literal},
	TokenData{LiteralStringBacktick, "Backtick", "sb", LiteralString},
	TokenData{LiteralStringChar, "Char", "sc", LiteralString},
	TokenData{LiteralStringDoc, "Doc", "sd", LiteralString},
	TokenData{LiteralStringDouble, "Double", "s2", LiteralString},
	TokenData{LiteralStringEscape, "Escape", "se", LiteralString},
	TokenData{LiteralStringHeredoc, "Heredoc", "sh", LiteralString},
	TokenData{LiteralStringInterpol, "Interpol", "si", LiteralString},
	TokenData{LiteralStringOther, "Other", "sx", LiteralString},
	TokenData{LiteralStringRegex, "Regex", "sr", LiteralString},
	TokenData{LiteralStringSingle, "Single", "s1", LiteralString},
	TokenData{LiteralStringSymbol, "Symbol", "ss", LiteralString},
	TokenData{LiteralNumber, "Number", "m", Literal},
	TokenData{LiteralNumberFloat, "Float", "mf", LiteralNumber},
	TokenData{LiteralNumberHex, "Hex", "mh", LiteralNumber},
	TokenData{LiteralNumberInteger, "Integer", "mi", LiteralNumber},
	TokenData{LiteralNumberIntegerLong, "Long", "il", LiteralNumberInteger},
	TokenData{LiteralNumberOct, "Oct", "mo", LiteralNumber},
	TokenData{LiteralNumberBin, "Bin", "mb", LiteralNumber},
	TokenData{LiteralNumberOther, "Other", "mx", LiteralNumber},

	TokenData{Operator, "Operator", "o", 0},
	TokenData{OperatorWord, "Word", "ow", Operator},

	TokenData{Punctuation, "Punctuation", "p", 0},
	TokenData{PunctuationIndicator, "PunctuationIndicator", "pi", Punctuation},

	TokenData{Comment, "Comment", "c", 0},
	TokenData{CommentDoc, "Doc", "cd", Comment},
	TokenData{CommentMultiline, "Multiline", "cm", Comment},
	TokenData{CommentPreproc, "Preproc", "cp", Comment},
	TokenData{CommentSingle, "Single", "c1", Comment},
	TokenData{CommentSpecial, "Special", "cs", Comment},

	TokenData{Generic, "Generic", "g", 0},
	TokenData{GenericDeleted, "Deleted", "gd", Generic},
	TokenData{GenericEmph, "Emph", "ge", Generic},
	TokenData{GenericError, "Error", "gr", Generic},
	TokenData{GenericHeading, "Heading", "gh", Generic},
	TokenData{GenericInserted, "Inserted", "gi", Generic},
	TokenData{GenericOutput, "Output", "go", Generic},
	TokenData{GenericPrompt, "Prompt", "gp", Generic},
	TokenData{GenericStrong, "Strong", "gs", Generic},
	TokenData{GenericSubheading, "Subheading", "gu", Generic},
	TokenData{GenericTraceback, "Traceback", "gt", Generic},
	TokenData{GenericLineno, "Lineno", "gl", Generic},
}

var TokenTable map[Token]TokenData

func init_token() {
	for _, td := range Tokens {
		TokenTable[td.Token] = td
	}
}

func ToName(t Token) (string, bool) {
	td, ok := TokenTable[t]
	if !ok {
		return "", false
	}
	return td.Name, true
}

func ToShortName(t Token) (string, bool) {
	td, ok := TokenTable[t]
	if !ok {
		return "", false
	}
	return td.ShortName, true
}

func ToParent(t Token) (Token, bool) {
	td, ok := TokenTable[t]
	if !ok {
		return 0, false
	}
	return td.Parent, true
}
