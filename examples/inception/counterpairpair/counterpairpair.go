// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package counterpairpair is an example using go-frp and two counterpairs.
package counterpairpair

import (
	"math/rand"

	"github.com/gmlewis/go-frp/examples/inception/counterpair"
	h "github.com/gmlewis/go-frp/html"
)

const max = 100

// MODEL

type Model struct {
	first counterpair.Model
	last  counterpair.Model
}

func Init(firstTop, firstBottom, lastTop, lastBottom int) Model {
	return Model{
		first: counterpair.Init(firstTop, firstBottom),
		last:  counterpair.Init(lastTop, lastBottom),
	}
}

// UPDATE

type Action func(Model) Model

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func ResetAll(model Model) Model { return Init(0, 0, 0, 0) }

func RandomizeAll(model Model) Model {
	return Init(rand.Intn(max), rand.Intn(max), rand.Intn(max), rand.Intn(max))
}

type CounterPairAction func(counterpair.Action) Model

func First(model Model) CounterPairAction {
	return func(action counterpair.Action) Model {
		return Model{
			first: model.first.Update(action),
			last:  model.last,
		}
	}
}

func Last(model Model) CounterPairAction {
	return func(action counterpair.Action) Model {
		return Model{
			first: model.first,
			last:  model.last.Update(action),
		}
	}
}

type WrapFunc func(model Model) interface{}

func identity(model Model) interface{} {
	return model
}

func wrapper(model Model, wrapFunc WrapFunc) func(action Action) interface{} {
	return func(action Action) interface{} {
		newModel := model.Update(action)
		return wrapFunc(newModel)
	}
}

func firstWrapper(model Model, wrapFunc WrapFunc) counterpair.WrapFunc {
	return func(cm counterpair.Model) interface{} {
		newModel := Model{
			first: cm,
			last:  model.last,
		}
		return wrapFunc(newModel)
	}
}

func lastWrapper(model Model, wrapFunc WrapFunc) counterpair.WrapFunc {
	return func(cm counterpair.Model) interface{} {
		newModel := Model{
			first: model.first,
			last:  cm,
		}
		return wrapFunc(newModel)
	}
}

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	var wf WrapFunc
	if wrapFunc == nil {
		wf = identity
	}
	return h.Div(
		m.first.View(rootUpdateFunc, firstWrapper(m, wf)),
		m.last.View(rootUpdateFunc, lastWrapper(m, wf)),
		h.Button(h.Text("Reset All")).OnClick(rootUpdateFunc, wrapper(m, wf), ResetAll),
		h.Button(h.Text("Randomize All")).OnClick(rootUpdateFunc, wrapper(m, wf), RandomizeAll),
	)
}
