package main

import (
	"github.com/gmlewis/go-frp/examples/1/counter"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := counter.Model(0)
	start.Start(m, counter.Updater(m))
}
