package reverse

import "testing"

func TestView(t *testing.T) {
	m := Model(0)
	v := m.View()
	want := `<div><input></input><label></label></div>`
	if got := v.String(); got != want {
		t.Errorf("View = %q, want %q", got, want)
	}
}
