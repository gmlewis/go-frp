// Package counterpair is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/2/CounterPair.elm
package counterpair

import (
	"github.com/gmlewis/go-frp/examples/2/counter"
	h "github.com/gmlewis/go-frp/html"
)

// MODEL

type Model struct {
	top    counter.Model
	bottom counter.Model
}

func Init(top, bottom int) Model {
	return Model{
		top:    counter.Init(top),
		bottom: counter.Init(bottom),
	}
}

// UPDATE

type Action func(Model) Model

func (m Model) Update(action Action) Model { return action(m) }

func Reset(model Model) Model { return Init(0, 0) }
func Top(action counter.Action, model Model) Model {
	return Model{
		top:    model.top.Update(action),
		bottom: model.bottom,
	}
}
func Bottom(action counter.Action, model Model) Model {
	return Model{
		top:    model.top,
		bottom: model.bottom.Update(action),
	}
}

// VIEW

func (m Model) View() h.HTML {
	return h.Div(
		m.top.View( /* address Signal.ForwardTo address Top */ ),
		m.bottom.View( /* address Signal.ForwardTo address Bottom */ ),
		h.Button(h.Text("Reset")).OnClick(m, Reset),
	)
}
