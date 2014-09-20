package beni

import (
	"io"
)

type Emitter func(Token, string) error

type Lexer interface {
	GetInfo() LexerInfo
	Parse(r io.Reader, emit Emitter) error
}

type LexerInfo struct {
	Name           string
	Aliases        []string
	Filenames      []string
	AliasFilenames []string
	Mimetypes      []string
	Priority       int
	Description    string
}

var Lexers = []Lexer{
	JavaLexer,
}
