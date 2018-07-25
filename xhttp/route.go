// Copyright (c) 2018 Claudemiro Alves Feitosa Neto
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
	// Priority The bigger, the sooner the route will be registered.
	// Important: The order of Routes with the same priority is not guarantee.
	Priority int

	// Method HTTP Method which the route must accept
	Method string

	// Path is the URL path
	// e.g: /hello/{name}
	Path string

	// Middleware a list of middleware that must be applied,
	// Important: Middleware are applied in the given order.
	Middleware []Middleware

	// Handler the standard go HTTP handler
	Handler http.Handler
}

// String returns a string representation of the route
func (r *Route) String() string {
	return fmt.Sprintf("Method: %s, Path: %s,  Priority: %d", r.Method, r.Path, r.Priority)
}

// RouteMapping Necessary to register more than one Route
type RouteMapping struct {
	fx.Out

	Route *Route `group:"x_route_mappings"`
}

// RouteMappings group routes to be registered by the Server.
// It is populated by the container with all routes from the group `routes`
type RouteMappings struct {
	fx.In

	Routes []*Route `group:"x_route_mappings"`
}

type NotFound struct {
	fx.In

	Handler http.Handler `name:"x_not_found_route_mapping" optional:"true"`
}

type MethodNotAllowed struct {
	fx.In

	Handler http.Handler `name:"x_method_not_allowed_route_mapping" optional:"true"`
}

type NotFoundRouteMapping struct {
	fx.Out

	Handler http.Handler `name:"x_not_found_route_mapping" optional:"true"`
}

type MethodNotAllowedRouteMapping struct {
	fx.Out

	Handler http.Handler `name:"x_method_not_allowed_route_mapping" optional:"true"`
}
