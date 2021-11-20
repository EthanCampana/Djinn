package colormanager

import (
	"djinn/token"
	"fmt"

	"github.com/fatih/color"
)

type ColorManager struct {
	colormap  map[token.TokenType]*color.Color
	c_default *color.Color
}

// This will have to get refactored at some point but idk when yet
func New() *ColorManager {
	m := map[token.TokenType]*color.Color{
		token.ASSIGN:   color.New(color.FgHiMagenta),
		token.PLUS:     color.New(color.FgHiMagenta),
		token.MINUS:    color.New(color.FgHiMagenta),
		token.ASTERISK: color.New(color.FgHiMagenta),
		token.LT:       color.New(color.FgHiMagenta),
		token.GT:       color.New(color.FgHiMagenta),
		token.EQ:       color.New(color.FgHiMagenta),
		token.IDENT:    color.New(color.FgHiWhite),
		token.BANG:     color.New(color.FgHiRed),
		token.NOT_EQ:   color.New(color.FgHiRed),
		token.LBRACE:   color.New(color.FgHiYellow),
		token.RBRACE:   color.New(color.FgHiYellow),
		token.LPAREN:   color.New(color.FgHiCyan),
		token.RPAREN:   color.New(color.FgHiCyan),
		token.TRUE:     color.New(color.FgGreen),
		token.FALSE:    color.New(color.FgRed),
		token.IF:       color.New(color.FgHiBlue),
		token.ELSE:     color.New(color.FgRed),
		token.CREATE:   color.New(color.FgYellow),
	}

	cm := &ColorManager{colormap: m, c_default: color.New(color.FgWhite)}
	return cm
}

func (cm *ColorManager) SetDefaultColor(c *color.Color) {
	cm.c_default = c
}

func (cm *ColorManager) SetTokenColor(t token.TokenType, c *color.Color) {
	cm.colormap[t] = c
}

func (cm *ColorManager) GenerateColor(x []token.Token) string {
	line := ""
	for _, s := range x {
		if color, ok := cm.colormap[s.Type]; ok {
			line += fmt.Sprintf("%s ", color.SprintFunc()(s.Literal))
		} else {
			line += fmt.Sprintf("%s ", cm.c_default.SprintFunc()(s.Literal))
		}

	}
	line += "\n"
	return line
}
