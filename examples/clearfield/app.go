package main

import (
	"github.com/gmlewis/go-frp/examples/clearfield/clearfield"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := clearfield.Model("Hello world!")
	start.Start(m, clearfield.Updater(m))
}
