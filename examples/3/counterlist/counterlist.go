package counterlist

import (
	c "github.com/gmlewis/go-frp/examples/3/counter"
	h "github.com/gmlewis/go-frp/html"
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

func Remove() {}
func Insert(model Model) Model {
	return Model{
		counters: append(model.counters, Counter{id: model.nextID, counter: c.Model(0)}),
		nextID:   model.nextID + 1,
	}
}

// VIEW

func (m Model) View() h.HTML {
	remove := h.Button(h.Text("Reset")).OnClick(m, Remove)
	insert := h.Button(h.Text("Add")).OnClick(m, Insert)
	p := []h.HTML{remove, insert}
	for _, counter := range m.counters {
		p = append(p, counter.counter.View( /* address Signal.ForwardTo address counter */ ))
	}
	return h.Div(p...)
}
