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

package xhealth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dimiro1/x/xhttp"
	"github.com/dimiro1/x/xlog"
	"github.com/dimiro1/x/xutils"

	"github.com/dimiro1/health"
	"go.uber.org/fx"
)

// Checker holds the name and a Checker function to be executed.
type Checker struct {
	Name    string
	Checker health.Checker
}

// CheckMapping wrap the return of a health check.
type CheckMapping struct {
	fx.Out

	Checker *Checker `group:"x_health"`
}

// HealthHandlerQualifier is necessary to give a name to the health check
type HealthHandlerQualifier struct {
	fx.Out

	Handler health.Handler `name:"x_healthcheck"`
}

// HealthHandler ir necessary to access the healthcheck by name
type HealthHandler struct {
	fx.In

	Handler health.Handler `name:"x_healthcheck"`
}

// Config holds health check configuration.
type Config struct {
	Path string
}

// LoadConfig create a new *Config and populate it with values from environment.
func LoadConfig() *Config {
	path := xutils.GetenvDefault("X_HEALTH_PATH", "/status")

	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}

	return &Config{
		Path: path,
	}
}

// ChecksMappings hold all checks registered by the container.
type ChecksMappings struct {
	fx.In

	Checks []*Checker `group:"x_health"`
}

// ProvideRouteMapping provide and Server HTTP Route to be registered by the server module.
func ProvideRouteMapping(c HealthHandler, cfg *Config) xhttp.RouteMapping {
	return xhttp.RouteMapping{
		Route: &xhttp.Route{
			Path:    cfg.Path,
			Method:  http.MethodGet,
			Handler: c.Handler,
		},
	}
}

// RegisterHealthChecks register the checks populated in ChecksMappings.
func RegisterHealthChecks(h HealthHandler, checks ChecksMappings, logger xlog.OptionalLogger) {
	if xlog.IsProvided(logger) {
		logger.Logger.Println("registering health checks")
	}

	for _, c := range checks.Checks {
		if xlog.IsProvided(logger) {
			logger.Logger.Printf("registering health %s", c.Name)
		}

		h.Handler.AddChecker(c.Name, c.Checker)
	}

	if xlog.IsProvided(logger) {
		logger.Logger.Println("finished registering health checks")
	}
}

// NewHealth create a new healthChecker and register the available checks.
func NewHealth() HealthHandlerQualifier {
	return HealthHandlerQualifier{
		Handler: health.NewHandler(),
	}
}
