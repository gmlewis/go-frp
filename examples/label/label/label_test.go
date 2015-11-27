package label

import (
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	m := Model("test string")
	v := m.View(Updater(m), nil)
	want := []string{
		"<div>", "</div>",
		`<input value="test string" id="0"/>`,
		"<label>test string</label>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}
