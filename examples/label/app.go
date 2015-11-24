package main

import (
	"github.com/gmlewis/go-frp/examples/label/label"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	start.Start(label.Model("Hello world!"))
}
