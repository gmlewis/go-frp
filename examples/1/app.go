// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gmlewis/go-frp/examples/1/counter"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := counter.Model(0)
	start.Start(m, counter.Updater(m))
}
