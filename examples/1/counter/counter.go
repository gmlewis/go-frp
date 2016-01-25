// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package counter is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/1/Counter.elm
package counter

import (
	"fmt"

	h "github.com/gmlewis/go-frp/html"
)

// MODEL

type Model int

func (m Model) String() string { return fmt.Sprintf("%v", int(m)) }

// UPDATE

type Action func(Model) Model

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func Increment(model Model) Model { return model + 1 }
func Decrement(model Model) Model { return model - 1 }

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	style := [][]string{
		{"font-size", "20px"},
		{"font-family", "monospace"},
		{"display", "inline-block"},
		{"width", "50px"},
		{"text-align", "center"},
	}
	return h.Div(
		h.Button(h.Text("-")).OnClick(rootUpdateFunc, Updater(m), Decrement),
		h.Div(h.Text(m.String())).Style(style),
		h.Button(h.Text("+")).OnClick(rootUpdateFunc, Updater(m), Increment),
	)
}
