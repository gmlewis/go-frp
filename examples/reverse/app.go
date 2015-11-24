package main

import (
	"github.com/gmlewis/go-frp/examples/reverse/reverse"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	start.Start(reverse.Model("Hello world!"))
}
