// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package clearfield

import (
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	m := Model("")
	v := m.View(Updater(m), nil)
	want := []string{
		"<div>", "</div>",
		`<input value=""/>`,
		"Clear</button>",
	}
	got, _ := v.Render()
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("View = %q, want %q", got, w)
		}
	}
}
