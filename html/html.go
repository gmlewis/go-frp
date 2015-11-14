// Package html provides functions to build up the DOM of an app.
package html

import "fmt"

// type Action func(App) App
//
// type App interface {
// 	Update(action Action) App
// 	View(address Address) HTML
// }

// Address is a place-holder for future signal handling.
type Address string

// HTML defines an HTML element.
type HTML struct {
	tag     string
	props   [][]string
	styles  [][]string
	body    string
	elems   []HTML
	address Address
	onClick interface{}
}

// String renders the HTML element into its string representation.
func (s HTML) String() string {
	var result string
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
		// onClick
		result += ">"
	}
	for _, v := range s.elems {
		result += v.String()
	}
	result += s.body
	if s.tag != "" {
		result += "</" + s.tag + ">"
	}
	return result
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

// OnClick adds an onClick handler to an HTML element.
func (s HTML) OnClick(address Address, action interface{}) HTML {
	s.address = address
	s.onClick = action
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
