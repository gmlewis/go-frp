// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package counterpair is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/2/CounterPair.elm
package counterpair

import (
	"github.com/gmlewis/go-frp/v2/examples/2/counter"
	h "github.com/gmlewis/go-frp/v2/html"
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

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func Reset(model Model) Model { return Init(0, 0) }

func top(model Model) counter.WrapFunc {
	return func(cm counter.Model) interface{} {
		return Model{
			top:    cm,
			bottom: model.bottom,
		}
	}
}

func bottom(model Model) counter.WrapFunc {
	return func(cm counter.Model) interface{} {
		return Model{
			top:    model.top,
			bottom: cm,
		}
	}
}

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	return h.Div(
		m.top.View(rootUpdateFunc, top(m)),
		m.bottom.View(rootUpdateFunc, bottom(m)),
		h.Button(h.Text("Reset")).OnClick(rootUpdateFunc, rootUpdateFunc, Reset),
	)
}
