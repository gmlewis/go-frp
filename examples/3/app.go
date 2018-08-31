// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gmlewis/go-frp/v2/examples/3/counterlist"
	"github.com/gmlewis/go-frp/v2/start"
)

func main() {
	m := counterlist.Init(0, 0, 0)
	start.Start(m, counterlist.Updater(m))
}
