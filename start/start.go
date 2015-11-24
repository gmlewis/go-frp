// Package start provides a simple framework for starting a go-frp web app.
package start

import (
	h "github.com/gmlewis/go-frp/html"
	"github.com/gopherjs/gopherjs/js"
)

// Action take a model, performs an action, and returns a new model.
type Action func(Model) Model

// Model requires that an Update and a View function be supplied.
// TODO(gmlewis): figure out how to handle the Update in a general type-safe way.
type Model interface {
	// Update(action Action) App
	View() h.HTML
}

// Start starts the web application.
// TODO(gmlewis): Support event handling and signals.
func Start(model Model) {
	v := model.View()
	str, initFuncs := v.Render()
	js.Global.Get("document").Call("write", str)
	for _, initFunc := range initFuncs {
		initFunc()
	}
}