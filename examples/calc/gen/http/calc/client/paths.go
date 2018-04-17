// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package client

import "fmt"

// AddCalcPath returns the URL path to the calc service add HTTP endpoint.
func AddCalcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}
