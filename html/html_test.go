// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package html

import (
	"testing"
)

func TestID(t *testing.T) {
	tests := []struct {
		el   HTML
		want string
	}{
		{el: HTML{id: ""}, want: "0"},
		{el: HTML{id: ""}, want: "1"}, // auto-increment
		{el: HTML{id: "user-defined-cached"}, want: "user-defined-cached"},
		{el: HTML{props: [][]string{{"yo", "ho"}, {"id", "user-defined-not-cached"}}}, want: "user-defined-not-cached"},
	}
	for _, test := range tests {
		if got := test.el.ID(); got != test.want {
			t.Errorf("ID %#v = %q, want %q", test, got, test.want)
		}
	}
}
