package lexer

import (
	"io"
	"io/ioutil"

	"github.com/koron/beni/token"
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
	ParseString(s string, e Emitter) error
	GetDebug() bool
	SetDebug(v bool)
}

type Factory interface {
	Info() Info
	New() (Lexer, error)
}

func Parse(l Lexer, r io.Reader, e Emitter) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return l.ParseString(string(b), e)
}
