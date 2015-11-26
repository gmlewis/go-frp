package main

import (
	"github.com/gmlewis/go-frp/examples/inception/counterpairpair"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	m := counterpairpair.Init(1, 2, 3, 4)
	start.Start(m, counterpairpair.Updater(m))
}
