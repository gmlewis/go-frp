// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gmlewis/go-frp/v2/examples/label/label"
	"github.com/gmlewis/go-frp/v2/start"
)

func main() {
	m := label.Model("Hello world!")
	start.Start(m, label.Updater(m))
}
