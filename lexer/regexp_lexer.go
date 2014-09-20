package lexer

import (
	"github.com/koron/beni/token"
	"io"
)

type RegexpLexerState int32

type RegexpLexerContext interface {
	Emit(t token.Code, s string) error
	Push(s RegexpLexerState) error
	Pop() error
	Delegate(l Lexer, s string) error
}

type RegexpBehavior func(c RegexpLexerContext, groups []string) error

type RegexpLexerRule struct {
	Pattern  string
	Behavior RegexpBehavior
}

type RegexpLexerData struct {
	Info
	States map[RegexpLexerState][]RegexpLexerRule
}

const (
	Root RegexpLexerState = iota

	JavaClass
	JavaImport
)

func RegexpEmit(t token.Code) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		return c.Emit(t, groups[0])
	}
}

func RegexpEmitPush(t token.Code, s RegexpLexerState) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Push(s)
	}
}

func RegexpEmitPop(t token.Code) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Pop()
	}
}

type RegexpLexer struct {
	Info Info
}

func NewRegexpLexer(d *RegexpLexerData) *RegexpLexer {
	return &RegexpLexer{
		Info: d.Info,
	}
}

func (l *RegexpLexer) GetInfo() Info {
	return l.Info
}

func (l *RegexpLexer) Parse(r io.Reader, e Emitter) error {
	// TODO:
	return nil
}
