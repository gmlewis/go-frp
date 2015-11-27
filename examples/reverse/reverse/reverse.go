// Package reverse is an example using go-frp modeled after the example found in
// the book "Functional Reactive Programming" by Stephen Blackheath and Anthony Jones:
// http://www.manning.com/books/functional-reactive-programming
package reverse

import (
	h "github.com/gmlewis/go-frp/html"
	"honnef.co/go/js/dom"
)

// MODEL

type Model string

func (m Model) String() string { return string(m) }

// UPDATE

type Action func(Model, dom.Event) Model

func Updater(model Model) func(Action, dom.Event) Model {
	return func(action Action, event dom.Event) Model { return model.Update(action, event) }
}
func (m Model) Update(action Action, event dom.Event) Model { return action(m, event) }

func Keypress(model Model, event dom.Event) Model {
	if t, ok := event.Target().(*dom.HTMLInputElement); ok {
		// TODO: save/restore focus and cursor position
		return Model(t.Value)
	}
	return model
}

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	return h.Div(
		h.Input(string(m)).OnKeypress(rootUpdateFunc, Updater(m), Keypress),
		h.Label(reverse(string(m))),
	)
}

// Reverse reverses a string. See:
// http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
