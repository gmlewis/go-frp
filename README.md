# go-frp - Functional Reactive Programming in Go

## Status
[![Build Status](https://travis-ci.org/gmlewis/go-frp.png)](https://travis-ci.org/gmlewis/go-frp)

## Overview

Functional Reactive Programming as demonstrated by [React][react],
[React Native][react-native] and [Flux][flux] appear to be game-changing
technologies for web development as well as native app development for both
Android and iOS.

[react]: http://facebook.github.io/react
[react-native]: http://facebook.github.io/react-native
[flux]: http://facebook.github.io/flux

While investigating these technologies, I came across the
[Elm Programming Language][elm] and was impressed by the simple web app model
and the idea that a single language could provide all the HTML, CSS, *and*
JavaScript necessary to create a full web application. I read through the
tutorials and examples and although I'm a big fan of functional programming
languages, it dawned on me that with the advent of [GopherJS][gopherjs], I could
do all this in Go.

[elm]: http://elm-lang.org/
[gopherjs]: https://github.com/gopherjs/gopherjs

This is not an official Google product.

## Why Go?

There are a number of reasons for using [Go][]:

* Go is easy to read and understand,
* [goimports][] (and its integration with text editors) is fantastic,
* if your code compiles, chances are good that it is correct, and
* Go makes programming fun again.

[Go]: http://www.golang.org/
[gofmt]: https://golang.org/cmd/gofmt
[goimports]: https://godoc.org/golang.org/x/tools/cmd/goimports

## Installation

```bash
$ go get -u github.com/gopherjs/gopherjs
$ go get -u github.com/gmlewis/go-frp
```

## Getting started

### Run the examples

```bash
$ cd $GOPATH/src/github.com/gmlewis/go-frp/examples/1
$ gopherjs build -m app.go
```
