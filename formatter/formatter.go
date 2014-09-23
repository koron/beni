package formatter

import (
	"github.com/koron/beni/theme"
	"github.com/koron/beni/token"
)

// Info is meta information of formatter.
type Info struct {
	Name      string
	Aliases   []string
	Filenames []string
}

// Factory for formatter.
type Factory interface {
	Info() Info
	New(t theme.Theme) (Formatter, error)
}

// Formatter format token stream.
type Formatter interface {
	Info() Info
	Format(c token.Code, s string) error
}

type formatter struct {
	info  Info
	theme theme.Theme
}

func (f *formatter) Info() Info {
	return f.info
}

func (f *formatter) lookup(c token.Code) theme.Style {
	return f.theme.GetStyle(c)
}
