package main

import (
	"github.com/gmlewis/go-frp/examples/label/label"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := label.Model("Hello world!")
	start.Start(m, label.Updater(m))
}
