// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package html provides functions to build up the DOM of an app.
package html

import (
	"fmt"
	"html/template"
	"log"
	"reflect"
	"strconv"

	"honnef.co/go/js/dom"
)

// domNode is the ID of the root <div> to write the generated DOM.
const domNode = "go-frp"

var nextID int // Keep all HTML IDs unique.

// A Viewer implements a View that renders to HTML.
type Viewer interface {
	View(rootUpdateFunc, updateFunc interface{}) HTML
}

// Render renders a view and writes it to the DOM.
func Render(view Viewer, rootUpdateFunc interface{}) {
	// log.Printf("GML: html.Render(view=%#v, rootUpdateFunc=%#v)", view, rootUpdateFunc)
	v := view.View(rootUpdateFunc, nil)
	str, initFuncs := v.Render()
	dom.GetWindow().Document().GetElementByID(domNode).SetInnerHTML(str)
	for _, initFunc := range initFuncs {
		initFunc()
	}
}

// HTML defines an HTML element.
type HTML struct {
	tag       string
	id        string
	props     [][]string
	styles    [][]string
	body      string
	elems     []HTML
	initFuncs []func()
}

// Render renders the HTML element into its string representation.
// It also surfaces all initFuncs to the top and returns them.
func (s HTML) Render() (string, []func()) {
	var result string
	var initFuncs []func()
	if s.tag != "" {
		result = "<" + s.tag
		for _, v := range s.props {
			result += fmt.Sprintf(" %s=%q", v[0], template.HTMLEscapeString(v[1]))
		}
		var styles string
		for _, v := range s.styles {
			styles += fmt.Sprintf("%s:%s;", v[0], template.HTMLEscapeString(v[1]))
		}
		if styles != "" {
			result += fmt.Sprintf(" style=%q", styles)
		}
		if s.body == "" && len(s.elems) == 0 {
			result += "/"
		}
		result += ">"
	}
	for _, v := range s.elems {
		str, ifs := v.Render()
		result += str
		initFuncs = append(initFuncs, ifs...)
	}
	initFuncs = append(initFuncs, s.initFuncs...)
	result += template.HTMLEscapeString(s.body)
	if s.tag != "" && !(s.body == "" && len(s.elems) == 0) {
		result += "</" + s.tag + ">"
	}
	return result, initFuncs
}

// Props adds a slice of properties to an HTML element.
func (s HTML) Props(props [][]string) HTML {
	s.props = props
	return s
}

// Style adds a slice of inline CSS styles to an HTML element.
func (s HTML) Style(styles [][]string) HTML {
	s.styles = styles
	return s
}

// ID returns the current HTML ID of this element (assigning if necessary).
func (s *HTML) ID() string {
	if s.id != "" {
		return s.id
	}
	// Check the properties to see if the user created an ID for the element.
	for _, p := range s.props {
		if p[0] == "id" || p[0] == "ID" {
			s.id = p[1]
			return s.id
		}
	}
	// Use the next-available numeric ID and bump it.
	s.id = strconv.Itoa(nextID)
	s.props = append(s.props, []string{"id", s.id})
	nextID++
	return s.id
}

// On adds an event handler to an HTML element.
func (s HTML) On(event string, rootUpdateFunc, updateFunc, actionFunc interface{}) HTML {
	id := s.ID()
	s.initFuncs = append(s.initFuncs, func() {
		el := dom.GetWindow().Document().GetElementByID(id)
		if el == nil {
			log.Printf("unable to find DOM element id=%q", id)
			return
		}
		el.AddEventListener(event, false, func(de dom.Event) {
			go func() {
				u := reflect.ValueOf(updateFunc)
				// log.Printf("GML: updateFunc=%#v, u=%#v", updateFunc, u)
				a := reflect.ValueOf(actionFunc)
				// log.Printf("GML: actionFunc=%#v, a=%#v, a.Type=%q", actionFunc, a, a.Type())
				args := []reflect.Value{a}
				if reflect.TypeOf(updateFunc).NumIn() == 2 {
					var e interface{} = de
					args = append(args, reflect.ValueOf(e))
				}
				newModel := u.Call(args)
				// log.Printf("GML: in click handler! len(newModel)=%v", len(newModel))
				// log.Printf("GML: in click handler! newModel[0]=%v", newModel[0])
				// log.Printf("GML: in click handler! newModel[0].Type=%v", newModel[0].Type())
				if view, ok := newModel[0].Interface().(Viewer); ok {
					Render(view, rootUpdateFunc)
				}
			}()
		})
	})
	return s
}

// OnClick adds an click handler to an HTML element.
func (s HTML) OnClick(rootUpdateFunc, updateFunc, actionFunc interface{}) HTML {
	return s.On("click", rootUpdateFunc, updateFunc, actionFunc)
}

// OnKeypress adds a keypress handler to an HTML element.
func (s HTML) OnKeypress(rootUpdateFunc, updateFunc, actionFunc interface{}) HTML {
	return s.On("keypress", rootUpdateFunc, updateFunc, actionFunc)
}

// Div creates an HTML <div>.
func Div(s ...HTML) HTML {
	return HTML{tag: "div", elems: s}
}

// Button creates an HTML <button>.
func Button(s HTML) HTML {
	return HTML{tag: "button", elems: []HTML{s}}
}

// Text creates an HTML text string.
// Quoting and XSS prevention will be added later.
func Text(s string) HTML {
	return HTML{body: s}
}

// Input creates an HTML <input>.
func Input(s string) HTML {
	return HTML{tag: "input", props: [][]string{{"value", s}}}
}

// Label creates an HTML <label>
func Label(s string) HTML {
	return HTML{tag: "label", body: s}
}
