package counterlist

import (
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	m := Init(1, 2, 3)
	v := m.View(Updater(m), nil)
	want := []string{
		"<div>", "</div>",
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">1</div>`,
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">2</div>`,
		`<div style="font-size:20px;font-family:monospace;display:inline-block;width:50px;text-align:center;">3</div>`,
		"-</button>",
		"+</button>",
		"Remove</button>",
		"Add</button>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}
