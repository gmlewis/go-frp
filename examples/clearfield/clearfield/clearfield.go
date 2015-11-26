// Package clearfield is an example using go-frp modeled after the example found in
// the book "Functional Reactive Programming" by Stephen Blackheath and Anthony Jones:
// http://www.manning.com/books/functional-reactive-programming
package clearfield

import h "github.com/gmlewis/go-frp/html"

// MODEL

type Model string

// UPDATE

type Action func(Model) Model

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func Clear(model Model) Model { return Model("") }

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	return h.Div(
		h.Input(string(m)),
		h.Button(h.Text("Clear")).OnClick(rootUpdateFunc, Updater(m), Clear),
	)
}
