package main

import (
	"github.com/gmlewis/go-frp/examples/2/counterpair"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := counterpair.Init(0, 0)
	start.Start(m, counterpair.Updater(m))
}
