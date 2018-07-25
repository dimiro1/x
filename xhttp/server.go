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

	"github.com/dimiro1/x/xlog"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type HTTPServer = *http.Server
type Router = *mux.Router

// Start starts the Server.
func Start(_ context.Context, server HTTPServer, logger xlog.OptionalLogger) error {
	if xlog.IsProvided(logger) {
		logger.Logger.Printf("starting server on %s", server.Addr)
	}
	return server.ListenAndServe()
}

// Stop stop the Server
func Stop(ctx context.Context, server HTTPServer, logger xlog.OptionalLogger) error {
	if xlog.IsProvided(logger) {
		logger.Logger.Printf("stopping server on %s", server.Addr)
	}
	return server.Shutdown(ctx)
}

// registerStartStop knows how to registerStartStop the server in the container lifecycle.
func registerStartStop(lc fx.Lifecycle, server HTTPServer, logger xlog.OptionalLogger) {
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
func NewEmptyRouter() Router {
	return mux.NewRouter()
}

func RegisterSafeHandlers(router Router, notFound NotFound, methodNotAllowed MethodNotAllowed, logger xlog.OptionalLogger) {
	if notFound.Handler != nil {
		if xlog.IsProvided(logger) {
			logger.Logger.Println("registering NotFoundHandler handler")
		}
		router.NotFoundHandler = notFound.Handler
	}

	if methodNotAllowed.Handler != nil {
		if xlog.IsProvided(logger) {
			logger.Logger.Println("registering MethodNotAllowedHandler handler")
		}
		router.MethodNotAllowedHandler = methodNotAllowed.Handler
	}
}

// RegisterRouteMappings register the given routes.
func RegisterRouteMappings(router Router, routes RouteMappings, logger xlog.OptionalLogger) {
	if xlog.IsProvided(logger) {
		logger.Logger.Println("registering routes")
	}

	for _, aRoute := range routes.Routes {
		if xlog.IsProvided(logger) {
			logger.Logger.Printf("registering %s", aRoute)
		}

		localRouter := router.NewRoute().Subrouter()

		// Applying middleware
		for _, m := range aRoute.Middleware {
			localRouter.Use(mux.MiddlewareFunc(m))
		}

		// Registering the handler
		localRouter.Handle(aRoute.Path, aRoute.Handler).Methods(aRoute.Method)
	}

	if xlog.IsProvided(logger) {
		logger.Logger.Println("finished registering routes")
	}
}

// NewHTTPServer returns a *http.Server with timeouts configured by *Config.
func NewHTTPServer(cfg *Config, router Router) HTTPServer {
	return &http.Server{
		Handler:      func() *mux.Router { return router }(),
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
