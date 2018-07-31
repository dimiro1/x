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
	"net/http"
)

//noinspection GoUnusedExportedFunction
func WithPriority(priority int) func(*Route) {
	return func(r *Route) {
		r.Priority = priority
	}
}

//noinspection GoUnusedExportedFunction
func WithMiddleware(middleware ...Middleware) func(*Route) {
	return func(r *Route) {
		r.Middleware = middleware
	}
}

func Handle(method, path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}

	for _, o := range options {
		o(route)
	}

	return RouteMapping{
		Route: route,
	}
}

func Get(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodGet, path, handler, options...)
}

func Post(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodPost, path, handler, options...)
}

func Put(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodPut, path, handler, options...)
}

func Delete(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodDelete, path, handler, options...)
}

func Options(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodOptions, path, handler, options...)
}

func Patch(path string, handler http.Handler, options ...func(*Route)) RouteMapping {
	return Handle(http.MethodPatch, path, handler, options...)
}
