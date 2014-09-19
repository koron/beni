package beni

import (
	"io"
	"regexp"
)

type State int32

const (
	Stay State = iota
	Pop
	Root

	Class
	Import
)

type Rule struct {
	Pattern string
	Token   Token
	Next    State

	regexp *regexp.Regexp
}

type RegexpLexer struct {
	Info   LexerInfo
	Tokens map[State][]*Rule
}

func (lex *RegexpLexer) GetInfo() LexerInfo {
	return lex.Info
}

func (lex *RegexpLexer) Parse(r io.Reader, emit Emitter) error {
	// TODO:
	return nil
}
