package main

import (
	"github.com/gmlewis/go-frp/examples/3/counterlist"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := counterlist.Init(0, 0, 0)
	start.Start(m, counterlist.Updater(m))
}
