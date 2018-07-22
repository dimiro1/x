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
	"context"
	"net/http"
	"time"

	"github.com/dimiro1/x/xlog"
	"github.com/dimiro1/x/xutils"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// Config holds server configuration
type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// LoadConfig create a new *Config and populate it with values from environment.
func LoadConfig() *Config {
	return &Config{
		Addr:         xutils.GetenvDefault("SERVER_ADDR", ":8080"),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 60,
	}
}

// RouteMappings group routes to be registered by the Server.
// It is populated by the container with all routes from the group `routes`
type RouteMappings struct {
	fx.In

	Routes []*Route `group:"x_route_mappings"`
}

// HTTPServerQualifier is necessary to give a name to the server
type HTTPServerQualifier struct {
	fx.Out

	Server *http.Server `name:"x_http_server"`
}

// HTTPServer ir necessary to access the server by name
type HTTPServer struct {
	fx.In

	Server *http.Server `name:"x_http_server"`
}

// Start starts the Server.
func Start(_ context.Context, server HTTPServer, logger xlog.Logger) error {
	logger.Logger.Printf("starting server on %s", server.Server.Addr)
	return server.Server.ListenAndServe()
}

// Stop stop the Server
func Stop(ctx context.Context, server HTTPServer, logger xlog.Logger) error {
	logger.Logger.Printf("stopping server on %s", server.Server.Addr)
	return server.Server.Shutdown(ctx)
}

// registerStartStop knows how to registerStartStop the server in the container lifecycle.
func registerStartStop(lc fx.Lifecycle, server HTTPServer, logger xlog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go Start(ctx, server, logger)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return Stop(ctx, server, logger)
		},
	})
}

// NewEmptyRouter creates a new HTTP server Mux and register the given routes.
func NewEmptyRouter() *mux.Router {
	return mux.NewRouter()
}

// RegisterRouteMappings register the given routes.
func RegisterRouteMappings(router *mux.Router, routes RouteMappings, logger xlog.Logger) {
	logger.Logger.Println("registering routes")
	for _, aRoute := range routes.Routes {
		logger.Logger.Printf("registering %s", aRoute)
		localRouter := router.NewRoute().Subrouter()

		// Applying middleware
		for _, m := range aRoute.Middleware {
			localRouter.Use(mux.MiddlewareFunc(m))
		}

		// Registering the handler
		localRouter.Handle(aRoute.Path, aRoute.Handler).Methods(aRoute.Method)
	}
	logger.Logger.Println("finished registering routes")
}

// NewHTTPServer returns a *http.Server with timeouts configured by *Config.
func NewHTTPServer(cfg *Config, mux *mux.Router) HTTPServerQualifier {
	return HTTPServerQualifier{
		Server: &http.Server{
			Handler:      mux,
			Addr:         cfg.Addr,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}
