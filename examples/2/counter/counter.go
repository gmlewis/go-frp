// Package counter is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/2/Counter.elm
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

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func Increment(model Model) Model { return model + 1 }
func Decrement(model Model) Model { return model - 1 }

type WrapFunc func(model Model) interface{}

func wrapper(model Model, wrapFunc WrapFunc) func(action Action) interface{} {
	return func(action Action) interface{} {
		newModel := model.Update(action)
		return wrapFunc(newModel)
	}
}

// VIEW

func (m Model) View(rootUpdateFunc interface{}, wrapFunc WrapFunc) h.HTML {
	style := [][]string{
		{"font-size", "20px"},
		{"font-family", "monospace"},
		{"display", "inline-block"},
		{"width", "50px"},
		{"text-align", "center"},
	}
	return h.Div(
		h.Button(h.Text("-")).OnClick(rootUpdateFunc, wrapper(m, wrapFunc), Decrement),
		h.Div(h.Text(m.String())).Style(style),
		h.Button(h.Text("+")).OnClick(rootUpdateFunc, wrapper(m, wrapFunc), Increment),
	)
}
