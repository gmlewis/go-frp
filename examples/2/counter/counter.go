package counter

import (
	"fmt"

	h "github.com/gmlewis/go-frp/html"
)

// MODEL

type Model int

func (m Model) String() string { return fmt.Sprintf("%v", int(m)) }
func Init(count int) Model     { return Model(count) }

// UPDATE

type Action func(Model) Model

func Increment(model Model) Model          { return model + 1 }
func Decrement(model Model) Model          { return model - 1 }
func (m Model) Update(action Action) Model { return action(m) }

// VIEW

func (m Model) View(address h.Address) h.HTML {
	style := [][]string{
		{"font-size", "20px"},
		{"font-family", "monospace"},
		{"display", "inline-block"},
		{"width", "50px"},
		{"text-align", "center"},
	}
	return h.Div(
		h.Button(h.Text("-")).OnClick(address, Decrement),
		h.Div(h.Text(m.String())).Style(style),
		h.Button(h.Text("+")).OnClick(address, Increment),
	)
}
