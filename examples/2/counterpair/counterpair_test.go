package counterpair

import (
	"strings"
	"testing"

	"github.com/gmlewis/go-frp/examples/2/counter"
)

func TestView(t *testing.T) {
	m := Init(1, 2)
	v := m.View()
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

func TestActions(t *testing.T) {
	tests := []struct {
		top, bottom             int
		count                   int
		actionTop, actionBottom counter.Action
		wantTop, wantBottom     int
	}{
		{10, 100, 5, counter.Increment, nil, 15, 100},
		{10, 100, 5, nil, counter.Increment, 10, 105},
		{10, 100, 5, counter.Increment, counter.Increment, 15, 105},
		{10, 100, 5, counter.Decrement, nil, 5, 100},
		{10, 100, 5, nil, counter.Decrement, 10, 95},
		{10, 100, 5, counter.Decrement, counter.Decrement, 5, 95},
		{10, 100, 5, counter.Increment, counter.Decrement, 15, 95},
		{10, 100, 5, counter.Decrement, counter.Increment, 5, 105},
	}
	for num, test := range tests {
		m := Init(test.top, test.bottom)
		for i := 0; i < test.count; i++ {
			if test.actionTop != nil {
				m = Top(test.actionTop, m)
			}
			if test.actionBottom != nil {
				m = Bottom(test.actionBottom, m)
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

func TestReset(t *testing.T) {
	m := Init(10, 100)
	m = m.Update(Reset)
	if got, want := int(m.top), 0; got != want {
		t.Errorf("Top: Reset = %v, want %v", got, want)
	}
	if got, want := int(m.bottom), 0; got != want {
		t.Errorf("Bottom: Reset = %v, want %v", got, want)
	}
}
