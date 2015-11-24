// Package html provides functions to build up the DOM of an app.
package html

import (
	"fmt"
	"log"
	"strconv"

	"honnef.co/go/js/dom"
)

var nextID int // Keep all HTML IDs unique.

// type Action func(App) App
//
// type App interface {
// 	Update(action Action) App
// 	View(address Address) HTML
// }

// Address is a place-holder for future signal handling.
// type Address string

// HTML defines an HTML element.
type HTML struct {
	tag    string
	id     string
	props  [][]string
	styles [][]string
	body   string
	elems  []HTML
	// address Address
	// onClick interface{}
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
			result += fmt.Sprintf(" %s=%q", v[0], v[1])
		}
		var styles string
		for _, v := range s.styles {
			styles += fmt.Sprintf("%s:%s;", v[0], v[1])
		}
		if styles != "" {
			result += fmt.Sprintf(" style=%q", styles)
		}
		result += ">"
	}
	for _, v := range s.elems {
		str, ifs := v.Render()
		result += str
		initFuncs = append(initFuncs, ifs...)
	}
	initFuncs = append(initFuncs, s.initFuncs...)
	result += s.body
	if s.tag != "" {
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

// OnClick adds an onClick handler to an HTML element.
func (s HTML) OnClick(model interface{}, action interface{}) HTML {
	// log.Printf("GML: OnClick... model=%#v, action=%#v", model, action)
	// s.address = address
	// s.onClick = action
	// s.props = append(s.props, []string{"onclick", "OnClickHandler()"})
	// use channels?
	// use an anonymous function?
	// js.Global.Get("myButton").Call("addEventListener", "click", func() { go func() {...}})
	id := s.ID()
	log.Printf("GML: creating element ID: s=%#v, model=%#v, id=%q", s, model, id)
	s.initFuncs = append(s.initFuncs, func() {
		log.Printf("GML: firing initFunc! s=%#v, model=%#v, id=%q", s, model, id)
		d := dom.GetWindow().Document()
		el := d.GetElementByID(id)
		el.AddEventListener("click", false, func(e dom.Event) {
			go func() {
				log.Printf("GML: in click handler! e=%#v, s=%#v, model=%#v, id=%q", e, s, model, id)
				// The following causes these errors: "Uncaught Error: reflect: call of ?FIXME? on func Value"
				// m := reflect.ValueOf(model)
				// log.Printf("GML: model=%#v, m=%#v, m.Type=%q, m.NumField=%v", model, m, m.Type(), m.NumField())
				// a := reflect.ValueOf(action)
				// log.Printf("GML: action=%#v, a=%#v, a.Type=%q, a.NumField=%v", action, a, a.Type(), a.NumField())
				// newModel := action(model)
				// log.Printf("GML: in click handler! newModel=%#v", newModel)
			}()
		})
	})
	return s
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
func Input() HTML {
	return HTML{tag: "input"}
}

// Label creates an HTML <label>
func Label() HTML {
	return HTML{tag: "label"}
}
