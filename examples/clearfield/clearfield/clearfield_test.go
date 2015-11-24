package clearfield

import (
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	m := Model(0)
	v := m.View()
	want := []string{
		"<div>", "</div>",
		"<input></input>",
		"Clear</button>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}
