package formatter

import (
	"github.com/koron/beni/theme"
	"github.com/koron/beni/token"
)

var term256Info = Info{
	Name:    "Terminal256",
	Aliases: []string{"terminal256", "console256", "256"},
}

type terminal256Factory struct {
}

func (*terminal256Factory) Info() Info {
	return term256Info
}

func (*terminal256Factory) New(t theme.Theme) (Formatter, error) {
	return &terminal256{
		info: term256Info,
	}, nil
}

type terminal256 struct {
	info Info
}

func (f *terminal256) Info() Info {
	return f.info
}

func (f *terminal256) Format(c token.Code, s string) error {
	// TODO:
	return nil
}

// Terminal256 formatter factory.
var Terminal256 = &terminal256Factory{}
