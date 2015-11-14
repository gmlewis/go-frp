package main

import (
	"github.com/gmlewis/go-frp/examples/2/counterpair"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	start.Start(counterpair.Init(0, 0))
}
