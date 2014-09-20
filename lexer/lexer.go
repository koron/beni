package lexer

import (
	"github.com/koron/beni/token"
	"io"
)

type Info struct {
	Name           string
	Aliases        []string
	Filenames      []string
	AliasFilenames []string
	Mimetypes      []string
	Priority       int
	Description    string
}

type Emitter interface {
	Emit(token.Code, string) error
}

type Lexer interface {
	GetInfo() Info
	Parse(r io.Reader, emit Emitter) error
}

type Factory interface {
	GetInfo() Info
	New() Lexer
}
