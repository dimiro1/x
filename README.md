# X

X is a set of modules that make creating HTTP servers fun again.

```go
package main

import (
	"io"
	"net/http"

	"github.com/dimiro1/health/url"
	"github.com/dimiro1/x"
	"github.com/dimiro1/x/xhealth"
	"github.com/dimiro1/x/xhttp"
	"github.com/dimiro1/x/xvars"
	"go.uber.org/fx"
)

type IndexMiddleware struct {
	fx.In

	Compress xhttp.Middleware `name:"hf_compress_middleware"`
}

// Handler using the compress middleware
func Index(indexMiddleware IndexMiddleware) xhttp.RouteMapping {
	return xhttp.RouteMapping{
		Route: &xhttp.Route{
			Path:   "/",
			Method: http.MethodGet,
			Middleware: []xhttp.Middleware{
				indexMiddleware.Compress,
			},
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, "Index Page")
			}),
		},
	}
}

func GoogleHealthCheck() xhealth.CheckMapping {
	return xhealth.CheckMapping{
		Checker: &xhealth.Checker{
			Name:    "google",
			Checker: url.NewChecker("http://www.google.com"),
		},
	}
}

func main() {
	fx.New(
		x.Module,
		xvars.Module,
		fx.Provide(
			Index,
			GoogleHealthCheck,
		),
	).Run()
}
```

# The name

`x` in math means multiplication, so, `x` in this library means that your productivity will be increased by multiple times.

# THANKS

For the https://github.com/uber-go to create and share the fantastic https://github.com/uber-go/fx package.

# Author

Claudemiro Alves Feitosa Neto

# LICENSE

Copyright 2018 Claudemiro Alves Feitosa Neto. All rights reserved.
Use of this source code is governed by a MIT-style
license that can be found in the LICENSE file.