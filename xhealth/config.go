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
	"strings"

	"go.uber.org/config"
	"go.uber.org/fx"
)

// Config holds health check configuration.
type Config struct {
	Path string
}

type LoadConfigParams struct {
	fx.In

	Provider config.Provider `optional:"true"`
}

func LoadConfig(params LoadConfigParams) (Config, error) {
	cfg := Config{
		Path: "/status",
	}

	if params.Provider != nil {
		err := params.Provider.Get("xhealth").Populate(&cfg)
		if err != nil {
			return cfg, err
		}
	}

	if !strings.HasPrefix(cfg.Path, "/") {
		cfg.Path = fmt.Sprintf("/%s", cfg.Path)
	}

	return cfg, nil
}
