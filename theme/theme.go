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

type Theme interface {
	GetName() string
	GetStyle(tc token.Code) Style
	GetColor(cc ColorCode) Color
}

type ThemeDefinition struct {
	Name     string
	Palettes map[ColorCode]Color
	Styles   map[token.Code]Style
}

func (t *ThemeDefinition) GetName() string {
	return t.Name
}

func (t *ThemeDefinition) GetStyle(tc token.Code) Style {
	// FIXME: style cascading
	s := t.findStyle(tc)
	for s != nil {
		tc = tc.Parent()
		if tc == 0 {
			return Style{}
		}
		s = t.findStyle(tc)
	}
	return *s
}

func (t *ThemeDefinition) GetColor(cc ColorCode) Color {
	color, ok := t.Palettes[cc]
	if !ok {
		return Color{}
	}
	return color
}

func (t *ThemeDefinition) findStyle(tc token.Code) *Style {
	s, ok := t.Styles[tc]
	if !ok {
		return nil
	}
	return &s
}
