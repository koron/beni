package main

import (
	"fmt"
	"github.com/koron/beni/lexer"
	"github.com/koron/beni/token"
	"os"
)

type emitter struct {
}

func (e *emitter) Emit(c token.Code, s string) error {
	fmt.Printf("%s: %q\n", c.Name(), s)
	return nil
}

func parse(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	l, err := lexer.Java.New()
	//l.SetDebug(true)
	if err != nil {
		return err
	}
	return lexer.Parse(l, f, &emitter{})
}

func main() {
	for _, name := range os.Args[1:] {
		if err := parse(name); err != nil {
			panic(err)
		}
	}
}
