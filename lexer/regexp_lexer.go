package lexer

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/koron/beni/token"
	"regexp"
	"strings"
	"unicode/utf8"
)

type RegexpLexerContext interface {
	Emit(t token.Code, s string) error
	Push(s RegexpLexerState) error
	Pop() error
	ParseString(s string) error
}

type RegexpAction func(c RegexpLexerContext, groups []string) error

// RegexpEmit generate "emit" action.
func RegexpEmit(t token.Code) RegexpAction {
	return func(c RegexpLexerContext, groups []string) error {
		return c.Emit(t, groups[0])
	}
}

// RegexpEmitPush generate "emit and push" action.
func RegexpEmitPush(t token.Code, s RegexpLexerState) RegexpAction {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Push(s)
	}
}

// RegexpEmitPush generate "emit and pop" action.
func RegexpEmitPop(t token.Code) RegexpAction {
	return func(c RegexpLexerContext, groups []string) error {
		if err := c.Emit(t, groups[0]); err != nil {
			return err
		}
		return c.Pop()
	}
}

type RegexpLexerRule struct {
	Pattern string
	Action  RegexpAction
}

func (r RegexpLexerRule) Convert() (*regexpRule, error) {
	rx, err := regexp.Compile(r.Pattern)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, r.Pattern)
	}
	return &regexpRule{regexp: rx, action: r.Action}, nil
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

type RegexpLexerDef struct {
	Info
	States map[RegexpLexerState][]RegexpLexerRule
}

type regexpRule struct {
	regexp *regexp.Regexp
	action RegexpAction
}

func (r *regexpRule) match(s string) []string {
	return r.regexp.FindStringSubmatch(s)
}

func (r *regexpRule) String() string {
	return fmt.Sprintf("rule{pattern=%s}", r.regexp.String())
}

type RegexpLexer struct {
	Info   Info
	States map[RegexpLexerState][]*regexpRule
	Debug  bool
}

func NewRegexpLexer(d *RegexpLexerDef) (*RegexpLexer, error) {
	states := make(map[RegexpLexerState][]*regexpRule)
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

func (l *RegexpLexer) ParseString(s string, e Emitter) error {
	c := &regexpLexerContext{
		lexer:      l,
		emitter:    e,
		stateStack: list.New(),
		debug:      l.Debug,
	}
	c.stateStack.PushBack(Root)
	return c.parse(s)
}

func (l *RegexpLexer) GetDebug() bool {
	return l.Debug
}

func (l *RegexpLexer) SetDebug(v bool) {
	l.Debug = v
}

func (l *RegexpLexer) Rules(s RegexpLexerState) []*regexpRule {
	rules, ok := l.States[s]
	if !ok {
		return nil
	}
	return rules
}

type regexpLexerContext struct {
	lexer      *RegexpLexer
	emitter    Emitter
	stateStack *list.List
	debug      bool
}

func (c *regexpLexerContext) Emit(t token.Code, s string) error {
	return c.emitter.Emit(t, s)
}

func (c *regexpLexerContext) Push(s RegexpLexerState) error {
	// FIXME: check many push.
	c.stateStack.PushBack(s)
	return nil
}

func (c *regexpLexerContext) Pop() error {
	if c.stateStack.Len() <= 0 {
		return errors.New("over pop")
	}
	e := c.stateStack.Back()
	c.stateStack.Remove(e)
	v, ok := e.Value.(RegexpLexerState)
	if !ok {
		return fmt.Errorf("unknown state: %v", v)
	}
	return nil
}

func (c *regexpLexerContext) ParseString(s string) error {
	prev := c.stateStack.Len()
	if err := c.parse(s); err != nil {
		return err
	}
	curr := c.stateStack.Len()
	switch {
	case curr < prev:
		return fmt.Errorf("over pop: %d < %d", curr, prev)
	case curr > prev:
		for curr > prev {
			c.stateStack.Remove(c.stateStack.Back())
			curr--
		}
	}
	return nil
}

func (c *regexpLexerContext) currentState() RegexpLexerState {
	v, _ := c.stateStack.Back().Value.(RegexpLexerState)
	return v
}

func (c *regexpLexerContext) debugf(s string, a ...interface{}) {
	if c.debug {
		fmt.Print("RegexpLexer:DEBUG:")
		fmt.Printf(s, a...)
	}
}

func (c *regexpLexerContext) parse(s string) error {
ParseLoop:
	for len(s) > 0 {
		rules := c.lexer.Rules(c.currentState())
		if rules == nil {
			return fmt.Errorf("unknown state: %v", c.currentState())
		}
		for _, rule := range rules {
			m := rule.match(s)
			if m == nil {
				continue
			}
			c.debugf("s=%v r=%s m=%q\n", c.currentState().String(), rule, m)
			if len(m[0]) == 0 {
				return errors.New("matched with empty")
			}
			if err := rule.action(c, m); err != nil {
				return err
			}
			s = s[len(m[0]):]
			continue ParseLoop
		}
		// forward pointer if no rules matched.
		ch, n := utf8.DecodeRuneInString(s)
		// FIXME: need to emit 'ch'?
		c.debugf("skip %q\n", ch)
		s = s[n:]
	}
	return nil
}

func regexpLexerJoin(words []string) string {
	quoted := make([]string, len(words))
	for i, v := range words {
		quoted[i] = regexp.QuoteMeta(v)
	}
	return strings.Join(quoted, "|")
}
