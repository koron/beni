package theme

import (
	"github.com/koron/beni/token"
)

const (
	base16_00 ColorCode = iota
	base16_01           = iota
	base16_02           = iota
	base16_03           = iota
	base16_04           = iota
	base16_05           = iota
	base16_06           = iota
	base16_07           = iota
	base16_08           = iota
	base16_09           = iota
	base16_0A           = iota
	base16_0B           = iota
	base16_0C           = iota
	base16_0D           = iota
	base16_0E           = iota
	base16_0F           = iota
)

var Base16 = &ThemeDefinition{
	Name: "base16",

	Palettes: map[ColorCode]Color{
		base16_00: Color{0x15, 0x15, 0x15},
		base16_01: Color{0x20, 0x20, 0x20},
		base16_02: Color{0x30, 0x30, 0x30},
		base16_03: Color{0x50, 0x50, 0x50},
		base16_04: Color{0xb0, 0xb0, 0xb0},
		base16_05: Color{0xd0, 0xd0, 0xd0},
		base16_06: Color{0xe0, 0xe0, 0xe0},
		base16_07: Color{0xf5, 0xf5, 0xf5},
		base16_08: Color{0xac, 0x41, 0x42},
		base16_09: Color{0xd2, 0x84, 0x45},
		base16_0A: Color{0xf4, 0xbf, 0x75},
		base16_0B: Color{0x90, 0xa9, 0x59},
		base16_0C: Color{0x75, 0xb5, 0xaa},
		base16_0D: Color{0x6a, 0x9f, 0xb5},
		base16_0E: Color{0xaa, 0x75, 0x9f},
		base16_0F: Color{0x8f, 0x55, 0x36},
	},

	Styles: map[token.Code]Style{
		token.Text:  Style{Fg: base16_05, Bg: base16_00},
		token.Error: Style{Fg: base16_00, Bg: base16_08},

		token.Comment:        Style{Fg: base16_03},
		token.CommentPreproc: Style{Fg: base16_0A},

		token.NameTag:     Style{Fg: base16_0A},
		token.Operator:    Style{Fg: base16_05},
		token.Punctuation: Style{Fg: base16_05},

		token.GenericInserted: Style{Fg: base16_0B},
		token.GenericDeleted:  Style{Fg: base16_08},
		token.GenericHeading: Style{
			Fg:   base16_0D,
			Bg:   base16_00,
			Bold: true,
		},

		token.Keyword:            Style{Fg: base16_0E},
		token.KeywordConstant:    Style{Fg: base16_09},
		token.KeywordType:        Style{Fg: base16_09},
		token.KeywordDeclaration: Style{Fg: base16_09},

		token.LiteralString:         Style{Fg: base16_0B},
		token.LiteralStringRegex:    Style{Fg: base16_0C},
		token.LiteralStringInterpol: Style{Fg: base16_0F},
		token.LiteralStringEscape:   Style{Fg: base16_0F},

		token.NameNamespace: Style{Fg: base16_0A},
		token.NameClass:     Style{Fg: base16_0A},
		token.NameConstant:  Style{Fg: base16_0A},
		token.NameAttribute: Style{Fg: base16_0D},

		token.LiteralNumber:       Style{Fg: base16_0B},
		token.LiteralStringSymbol: Style{Fg: base16_0B},
	},
}
