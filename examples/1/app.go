package main

import (
	"github.com/gmlewis/go-frp/examples/1/counter"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	start.Start(counter.Model(0))
}
