// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package counterlist is an example using go-frp modeled after the Elm example found at:
// https://github.com/evancz/elm-architecture-tutorial/blob/master/examples/3/CounterList.elm
package counterlist

import (
	c "github.com/gmlewis/go-frp/v2/examples/3/counter"
	h "github.com/gmlewis/go-frp/v2/html"
)

// MODEL

type Model struct {
	counters []Counter
	nextID   ID
}

type ID int

type Counter struct {
	id      ID
	counter c.Model
}

func Init(values ...int) Model {
	m := Model{nextID: ID(len(values))}
	for id, value := range values {
		m.counters = append(m.counters, Counter{id: ID(id), counter: c.Model(value)})
	}
	return m
}

// UPDATE

type Action func(Model) Model

func Updater(model Model) func(action Action) Model {
	return func(action Action) Model { return model.Update(action) }
}
func (m Model) Update(action Action) Model { return action(m) }

func Remove(model Model) Model {
	var counters []Counter
	var nextID ID
	if len(model.counters) > 1 {
		counters = model.counters[:len(model.counters)-1]
		nextID = model.nextID - 1
	}
	return Model{
		counters: counters,
		nextID:   nextID,
	}
}

func Insert(model Model) Model {
	counters := model.counters[:] // Copy counters
	return Model{
		counters: append(counters, Counter{id: model.nextID, counter: c.Model(0)}),
		nextID:   model.nextID + 1,
	}
}

func cWrapper(model Model, id int) c.WrapFunc {
	return func(cm c.Model) interface{} {
		var counters []Counter
		for i := 0; i < len(model.counters); i++ {
			if i == id {
				counters = append(counters, Counter{id: ID(id), counter: cm})
				continue
			}
			counters = append(counters, model.counters[i])
		}
		return Model{
			counters: counters,
			nextID:   model.nextID,
		}
	}
}

func wrapper(model Model) func(action Action) interface{} {
	return func(action Action) interface{} {
		return model.Update(action)
	}
}

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	remove := h.Button(h.Text("Remove")).OnClick(rootUpdateFunc, wrapper(m), Remove)
	insert := h.Button(h.Text("Add")).OnClick(rootUpdateFunc, wrapper(m), Insert)
	p := []h.HTML{remove, insert}
	for id, counter := range m.counters {
		p = append(p, counter.counter.View(rootUpdateFunc, cWrapper(m, id)))
	}
	return h.Div(p...)
}
