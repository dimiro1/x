// Copyright (x) 2018 Claudemiro Alves Feitosa Neto
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package xhttp

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

// Route holds everything the router needs to know to register a route into the container
type Route struct {
	// Method HTTP Method which the route must accept
	Method string

	// Path is the URL path
	// e.g: /hello/{name}
	Path string

	// Middleware a list of middleware that must be applied
	Middleware []Middleware

	// Handler the standard go HTTP handler
	Handler http.Handler
}

// String returns a string representation of the route
func (r *Route) String() string {
	return fmt.Sprintf("%s %s", r.Method, r.Path)
}

// RouteMapping Necessary to register more than one Route
type RouteMapping struct {
	fx.Out

	Route *Route `group:"hf_route_mappings"`
}