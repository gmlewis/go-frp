// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package counterpair

import (
	"strings"
	"testing"

	"github.com/gmlewis/go-frp/examples/inception/counter"
)

func TestView(t *testing.T) {
	m := Init(1, 2)
	v := m.View(Updater(m), nil)
	want := []string{
		"<div>", "</div>",
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">1</div>`,
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">2</div>`,
		"-</button>",
		"+</button>",
		"Reset</button>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}

func TestCounterActions(t *testing.T) {
	const (
		top    = 10
		bottom = 100
		count  = 5
	)
	tests := []struct {
		actionTop, actionBottom counter.Action
		wantTop, wantBottom     int
	}{
		{counter.Increment, nil, 15, 100},
		{nil, counter.Increment, 10, 105},
		{counter.Increment, counter.Increment, 15, 105},
		{counter.Decrement, nil, 5, 100},
		{nil, counter.Decrement, 10, 95},
		{counter.Decrement, counter.Decrement, 5, 95},
		{counter.Increment, counter.Decrement, 15, 95},
		{counter.Decrement, counter.Increment, 5, 105},
	}
	for num, test := range tests {
		m := Init(top, bottom)
		for i := 0; i < count; i++ {
			if test.actionTop != nil {
				m = Top(m)(test.actionTop)
			}
			if test.actionBottom != nil {
				m = Bottom(m)(test.actionBottom)
			}
		}
		if gotTop := int(m.top); gotTop != test.wantTop {
			t.Errorf("test #%v: Top = %v, want %v", num, gotTop, test.wantTop)
		}
		if gotBottom := int(m.bottom); gotBottom != test.wantBottom {
			t.Errorf("test #%v: Bottom = %v, want %v", num, gotBottom, test.wantBottom)
		}
	}
}

func TestActions(t *testing.T) {
	const (
		top    = 10
		bottom = 100
		count  = 5
	)
	tests := []struct {
		action              Action
		wantTop, wantBottom int
	}{
		{IncrementTop, 15, 100},
		{IncrementBottom, 10, 105},
		{AdjustBy(1, 1), 15, 105},
		{DecrementTop, 5, 100},
		{DecrementBottom, 10, 95},
		{AdjustBy(-1, -1), 5, 95},
		{AdjustBy(1, -1), 15, 95},
		{AdjustBy(-1, 1), 5, 105},
	}
	for num, test := range tests {
		m := Init(top, bottom)
		for i := 0; i < count; i++ {
			m = Updater(m)(test.action)
		}
		if gotTop := int(m.top); gotTop != test.wantTop {
			t.Errorf("test #%v: Top = %v, want %v", num, gotTop, test.wantTop)
		}
		if gotBottom := int(m.bottom); gotBottom != test.wantBottom {
			t.Errorf("test #%v: Bottom = %v, want %v", num, gotBottom, test.wantBottom)
		}
	}
}

func TestReset(t *testing.T) {
	m := Init(10, 100)
	m = Updater(m)(Reset)
	if got, want := int(m.top), 0; got != want {
		t.Errorf("Top: Reset = %v, want %v", got, want)
	}
	if got, want := int(m.bottom), 0; got != want {
		t.Errorf("Bottom: Reset = %v, want %v", got, want)
	}
}
