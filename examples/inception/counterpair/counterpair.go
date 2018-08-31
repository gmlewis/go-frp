// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package counterpair is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/inception/CounterPair.elm
package counterpair

import (
	"github.com/gmlewis/go-frp/v2/examples/inception/counter"
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

func (m Model) Top() int    { return int(m.top) }
func (m Model) Bottom() int { return int(m.bottom) }

// UPDATE

type Action func(Model) Model

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func IncrementTop(model Model) Model    { return Top(model)(counter.Increment) }
func IncrementBottom(model Model) Model { return Bottom(model)(counter.Increment) }
func DecrementTop(model Model) Model    { return Top(model)(counter.Decrement) }
func DecrementBottom(model Model) Model { return Bottom(model)(counter.Decrement) }
func Reset(model Model) Model           { return Init(0, 0) }

type CounterAction func(counter.Action) Model

func Top(model Model) CounterAction {
	return func(action counter.Action) Model {
		return Model{
			top:    model.top.Update(action),
			bottom: model.bottom,
		}
	}
}

func Bottom(model Model) CounterAction {
	return func(action counter.Action) Model {
		return Model{
			top:    model.top,
			bottom: model.bottom.Update(action),
		}
	}
}

func AdjustBy(top, bottom int) func(model Model) Model {
	return func(model Model) Model {
		return Model{
			top:    counter.Init(int(model.top) + top),
			bottom: counter.Init(int(model.bottom) + bottom),
		}
	}
}

type WrapFunc func(model Model) interface{}

func wrapper(model Model, wrapFunc WrapFunc) func(action Action) interface{} {
	return func(action Action) interface{} {
		newModel := model.Update(action)
		return wrapFunc(newModel)
	}
}

func topWrapper(model Model, wrapFunc WrapFunc) counter.WrapFunc {
	return func(cm counter.Model) interface{} {
		newModel := Model{
			top:    cm,
			bottom: model.bottom,
		}
		return wrapFunc(newModel)
	}
}

func bottomWrapper(model Model, wrapFunc WrapFunc) counter.WrapFunc {
	return func(cm counter.Model) interface{} {
		newModel := Model{
			top:    model.top,
			bottom: cm,
		}
		return wrapFunc(newModel)
	}
}

// VIEW

func (m Model) View(rootUpdateFunc interface{}, wrapFunc WrapFunc) h.HTML {
	return h.Div(
		m.top.View(rootUpdateFunc, topWrapper(m, wrapFunc)),
		m.bottom.View(rootUpdateFunc, bottomWrapper(m, wrapFunc)),
		h.Button(h.Text("Reset")).OnClick(rootUpdateFunc, wrapper(m, wrapFunc), Reset),
	)
}
