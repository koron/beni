package main

import (
	"fmt"
	"os"

	"github.com/koron/beni/formatter"
	"github.com/koron/beni/lexer"
	"github.com/koron/beni/theme"
	"github.com/koron/beni/token"
)

type emitter struct {
	formatter formatter.Formatter
}

func (e *emitter) Emit(c token.Code, s string) error {
	return e.formatter.Format(c, s)
}

func parse(name string) error {
	fmt.Println(name)
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	l, err := lexer.Go.New()
	//l.SetDebug(true)
	if err != nil {
		return err
	}

	t, err := formatter.Terminal256.New(theme.Base16, os.Stdout)
	if err != nil {
		return err
	}

	return lexer.Parse(l, f, &emitter{formatter: t})
}

func main() {
	for _, name := range os.Args[1:] {
		if err := parse(name); err != nil {
			panic(err)
		}
	}
}
