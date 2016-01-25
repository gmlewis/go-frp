// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gmlewis/go-frp/examples/clearfield/clearfield"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := clearfield.Model("Hello world!")
	start.Start(m, clearfield.Updater(m))
}
