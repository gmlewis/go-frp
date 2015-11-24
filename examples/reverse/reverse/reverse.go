// Package reverse is an example using go-frp modeled after the example found in
// the book "Functional Reactive Programming" by Stephen Blackheath and Anthony Jones:
// http://www.manning.com/books/functional-reactive-programming
package reverse

import h "github.com/gmlewis/go-frp/html"

// MODEL

type Model string

func (m Model) String() string { return string(m) }

// UPDATE

type Action func(Model) Model

func (m Model) Update(action Action) Model { return action(m) }

// VIEW

func (m Model) View() h.HTML {
	return h.Div(
		h.Input(),
		h.Label(),
	)
}
