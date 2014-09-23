package main

import (
	"fmt"
	"io"
	"os"

	"github.com/koron/beni/formatter"
	"github.com/koron/beni/lexer"
	"github.com/koron/beni/theme"
)

// HighlightOptions has highlight options.
type HighlightOptions struct {
	Lexer     string
	Theme     string
	Formatter string
}

func parse(name string, w io.Writer, lf lexer.Factory, t theme.Theme, ff formatter.Factory) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	l, err := lf.New()
	if err != nil {
		return err
	}
	f, err := ff.New(t, w)

	if err = f.Start(); err != nil {
		return err
	}
	if err = lexer.Parse(l, file, f); err != nil {
		return err
	}
	if err = f.End(); err != nil {
		return err
	}

	return nil
}

// Highlight decorates contents of a file with highlight.
func Highlight(filename string, w io.Writer, o HighlightOptions) error {
	l := lexer.Find(filename)
	if l == nil {
		return fmt.Errorf("not found lexer: %s", filename)
	}
	t := theme.Find(o.Theme)
	if t == nil {
		return fmt.Errorf("not found theme: %s", o.Theme)
	}
	f := formatter.Find(o.Formatter)
	if f == nil {
		return fmt.Errorf("not found formatter: %s", o.Formatter)
	}
	if err := parse(filename, w, l, t, f); err != nil {
		return err
	}
	return nil
}
