// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package start provides a simple framework for starting a go-frp web app.
package start

import h "github.com/gmlewis/go-frp/v2/html"

func Start(view h.Viewer, rootUpdateFunc interface{}) {
	// log.Printf("Start(view=%#v, rootUpdateFunc=%#v)", view, rootUpdateFunc)
	h.Render(view, rootUpdateFunc)
}
