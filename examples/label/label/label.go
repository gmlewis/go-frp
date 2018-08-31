// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package label is an example using go-frp modeled after the example found in
// the book "Functional Reactive Programming" by Stephen Blackheath and Anthony Jones:
// http://www.manning.com/books/functional-reactive-programming
package label

import (
	h "github.com/gmlewis/go-frp/v2/html"
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
		h.Label(string(m)),
	)
}
