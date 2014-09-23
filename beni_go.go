package main

import (
	"fmt"
	"os"

	"github.com/koron/beni/formatter"
	"github.com/koron/beni/lexer"
	"github.com/koron/beni/theme"
)

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

	if err = t.Start(); err != nil {
		return err
	}
	if err = lexer.Parse(l, f, t); err != nil {
		return err
	}
	if err = t.End(); err != nil {
		return err
	}

	return nil
}

func main() {
	for _, name := range os.Args[1:] {
		if err := parse(name); err != nil {
			panic(err)
		}
	}
}
