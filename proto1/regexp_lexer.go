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

type StateMachine interface {
	Emit(t Token, s string) error
	Push(t State) error
	Pop() error
	Delegate(l Lexer, s string) error
}

type Behavior func(m StateMachine, groups []string) error

type Rule struct {
	Pattern  string
	Token    Token
	Next     State
	Behavior Behavior

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
