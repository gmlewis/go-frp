package counterpairpair

import (
	"strings"
	"testing"

	"github.com/gmlewis/go-frp/examples/inception/counterpair"
)

func TestView(t *testing.T) {
	m := Init(1, 2, 3, 4)
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

func TestCounterPairActions(t *testing.T) {
	const (
		firstTop    = 10
		firstBottom = 100
		lastTop     = 50
		lastBottom  = 500
		count       = 5
	)
	tests := []struct {
		actionFirst, actionLast       counterpair.Action
		wantFirstTop, wantFirstBottom int
		wantLastTop, wantLastBottom   int
	}{
		{counterpair.IncrementTop, nil, 15, 100, 50, 500},
		{nil, counterpair.IncrementTop, 10, 100, 55, 500},
		{counterpair.IncrementTop, counterpair.IncrementTop, 15, 100, 55, 500},
		{counterpair.DecrementTop, nil, 5, 100, 50, 500},
		{nil, counterpair.DecrementTop, 10, 100, 45, 500},
		{counterpair.DecrementTop, counterpair.DecrementTop, 5, 100, 45, 500},
		{counterpair.IncrementTop, counterpair.DecrementTop, 15, 100, 45, 500},
		{counterpair.DecrementTop, counterpair.IncrementTop, 5, 100, 55, 500},
		{counterpair.IncrementBottom, nil, 10, 105, 50, 500},
		{nil, counterpair.IncrementBottom, 10, 100, 50, 505},
		{counterpair.IncrementBottom, counterpair.IncrementBottom, 10, 105, 50, 505},
		{counterpair.DecrementBottom, nil, 10, 95, 50, 500},
		{nil, counterpair.DecrementBottom, 10, 100, 50, 495},
		{counterpair.DecrementBottom, counterpair.DecrementBottom, 10, 95, 50, 495},
		{counterpair.IncrementBottom, counterpair.DecrementBottom, 10, 105, 50, 495},
		{counterpair.DecrementBottom, counterpair.IncrementBottom, 10, 95, 50, 505},
	}
	for num, test := range tests {
		m := Init(firstTop, firstBottom, lastTop, lastBottom)
		for i := 0; i < count; i++ {
			if test.actionFirst != nil {
				m = First(m)(test.actionFirst)
			}
			if test.actionLast != nil {
				m = Last(m)(test.actionLast)
			}
		}
		if gotFirstTop := int(m.first.Top()); gotFirstTop != test.wantFirstTop {
			t.Errorf("test #%v: FirstTop = %v, want %v", num, gotFirstTop, test.wantFirstTop)
		}
		if gotFirstBottom := int(m.first.Bottom()); gotFirstBottom != test.wantFirstBottom {
			t.Errorf("test #%v: FirstBottom = %v, want %v", num, gotFirstBottom, test.wantFirstBottom)
		}
		if gotLastTop := int(m.last.Top()); gotLastTop != test.wantLastTop {
			t.Errorf("test #%v: LastTop = %v, want %v", num, gotLastTop, test.wantLastTop)
		}
		if gotLastBottom := int(m.last.Bottom()); gotLastBottom != test.wantLastBottom {
			t.Errorf("test #%v: LastBottom = %v, want %v", num, gotLastBottom, test.wantLastBottom)
		}
	}
}

func TestResetAll(t *testing.T) {
	m := Init(10, 100, 50, 500)
	m = Updater(m)(ResetAll)
	if got, want := int(m.first.Top()), 0; got != want {
		t.Errorf("FirstTop: ResetAll = %v, want %v", got, want)
	}
	if got, want := int(m.first.Bottom()), 0; got != want {
		t.Errorf("FirstBottom: ResetAll = %v, want %v", got, want)
	}
	if got, want := int(m.last.Top()), 0; got != want {
		t.Errorf("LastTop: ResetAll = %v, want %v", got, want)
	}
	if got, want := int(m.last.Bottom()), 0; got != want {
		t.Errorf("LastBottom: ResetAll = %v, want %v", got, want)
	}
}
