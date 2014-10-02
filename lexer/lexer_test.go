package lexer

import (
	"testing"

	"github.com/koron/beni/token"
)

type Result struct {
	t       *testing.T
	Verbose bool
	Tokens  []interface{}
}

func (r *Result) Emit(c token.Code, s string) error {
	if r.Verbose {
		r.t.Logf("%s %#v", c.Name(), s)
	}
	r.Tokens = append(r.Tokens, c, s)
	return nil
}

func parseCheck(t *testing.T, f Factory, s string, tokens []interface{},
	verbose bool) {
	// Create lexer.
	l, err := f.New()
	if err != nil {
		t.Error(err)
	}
	// Parse.
	r := &Result{
		t:       t,
		Tokens:  make([]interface{}, 0, len(tokens)),
		Verbose: verbose,
	}
	if err := l.ParseString(s, r); err != nil {
		t.Error(err)
	}
	// Check tokens.
	len_exp := len(tokens)
	len_act := len(r.Tokens)
	if len_exp != len_act {
		t.Logf("not match length of tokens: expected=%d actual=%d",
			len_exp, len_act)
		t.Fail()
	}
	// Check contents of tokens.
	min := len_exp
	if min > len_act {
		min = len_act
	}
	for i := 0; i < min; i += 2 {
		var e0, e1 bool
		if tokens[i] != r.Tokens[i] {
			e0 = true
		}
		j := i + 1
		if j < min && tokens[j] != r.Tokens[j] {
			e1 = true
		}
		if e0 || e1 {
			if j < min {
				t.Errorf("not match at #%d: expected=(%s %#v) actual=(%s %#v)",
					i, tokens[i], tokens[j], r.Tokens[i], r.Tokens[j])
			} else {
				t.Errorf("not match at #%d: expected=(%s) actual=(%s)",
					i, tokens[i], r.Tokens[i])
			}
		}
	}
}
