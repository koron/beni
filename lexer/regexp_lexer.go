package lexer

import (
	"github.com/koron/beni/token"
	"io"
	"regexp"
)

type RegexpLexerContext interface {
	Emit(t token.Code, s string) error
	Push(s RegexpLexerState) error
	Pop() error
	ParseStr(str string) error
}

type RegexpBehavior func(c RegexpLexerContext, groups []string) error

// RegexpEmit generate "emit" behavior.
func RegexpEmit(t token.Code) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		return c.Emit(t, groups[0])
	}
}

// RegexpEmitPush generate "emit and push" behavior.
func RegexpEmitPush(t token.Code, s RegexpLexerState) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Push(s)
	}
}

// RegexpEmitPush generate "emit and pop" behavior.
func RegexpEmitPop(t token.Code) RegexpBehavior {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Pop()
	}
}

type RegexpLexerRule struct {
	Pattern  string
	Behavior RegexpBehavior
}

func (r RegexpLexerRule) Convert() (*regexpRule, error) {
	rx, err := regexp.Compile(r.Pattern)
	if err != nil {
		return nil, err
	}
	return &regexpRule{pattern: rx, behavior: r.Behavior}, nil
}

func regexpConvertRules(src []RegexpLexerRule) ([]*regexpRule, error) {
	dst := make([]*regexpRule, len(src))
	for i, rs := range src {
		rd, err := rs.Convert()
		if err != nil {
			return nil, err
		}
		dst[i] = rd
	}
	return dst, nil
}

type RegexpLexerDefinition struct {
	Info
	States map[RegexpLexerState][]RegexpLexerRule
}

type regexpRule struct {
	pattern  *regexp.Regexp
	behavior RegexpBehavior
}

type RegexpLexer struct {
	Info   Info
	States map[RegexpLexerState][]*regexpRule
}

func NewRegexpLexer(d *RegexpLexerDefinition) (*RegexpLexer, error) {
	var states map[RegexpLexerState][]*regexpRule
	for s, r := range d.States {
		rules, err := regexpConvertRules(r)
		if err != nil {
			return nil, err
		}
		states[s] = rules
	}
	return &RegexpLexer{Info: d.Info, States: states}, nil
}

func (l *RegexpLexer) GetInfo() Info {
	return l.Info
}

func (l *RegexpLexer) Parse(r io.Reader, e Emitter) error {
	// TODO:
	return nil
}
