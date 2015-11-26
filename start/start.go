// Package start provides a simple framework for starting a go-frp web app.
package start

import h "github.com/gmlewis/go-frp/html"

func Start(view h.Viewer, rootUpdateFunc interface{}) {
	// log.Printf("Start(view=%#v, rootUpdateFunc=%#v)", view, rootUpdateFunc)
	h.Render(view, rootUpdateFunc)
}
