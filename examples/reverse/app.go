package main

import (
	"github.com/gmlewis/go-frp/examples/reverse/reverse"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := reverse.Model("Hello world!")
	start.Start(m, reverse.Updater(m))
}
