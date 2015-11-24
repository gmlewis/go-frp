package counter

import (
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	m := Model(0)
	v := m.View()
	want := []string{
		"<div>", "</div>",
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">0</div>`,
		"-</button>",
		"+</button>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}

func TestIncrement(t *testing.T) {
	m := Model(0)
	for i := 0; i < 10; i++ {
		m = Increment(m)
	}
	if got, want := int(m), 10; got != want {
		t.Errorf("Increment = %v, want %v", got, want)
	}
}

func TestDecrement(t *testing.T) {
	m := Model(0)
	for i := 0; i < 10; i++ {
		m = Decrement(m)
	}
	if got, want := int(m), -10; got != want {
		t.Errorf("Decrement = %v, want %v", got, want)
	}
}

func TestUpdate(t *testing.T) {
	m := Model(0)
	m = m.Update(Increment)
	if got, want := int(m), 1; got != want {
		t.Errorf("Increment = %v, want %v", got, want)
	}
	m = m.Update(Decrement)
	if got, want := int(m), 0; got != want {
		t.Errorf("Decrement = %v, want %v", got, want)
	}
}
