package counterpair

import (
	"strings"
	"testing"
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
