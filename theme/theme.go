package theme

import (
	"github.com/koron/beni/token"
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type ColorCode int32

type Style struct {
	Fg   ColorCode
	Bg   ColorCode
	Bold bool
}

type ThemeDefinition struct {
	Name     string
	Palettes map[ColorCode]Color
	Styles   map[token.Code]Style
}
