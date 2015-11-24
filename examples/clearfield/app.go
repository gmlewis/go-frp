package main

import (
	"github.com/gmlewis/go-frp/examples/clearfield/clearfield"
	"github.com/gmlewis/go-frp/start"
)

func main() {
	start.Start(clearfield.Model("Hello world!"))
}
