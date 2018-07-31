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

package xlog

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

// Config holds log configuration
type Config struct {
	Prefix string `yaml:"prefix"`
}

type LoadConfigParams struct {
	fx.In

	Provider config.Provider `optional:"true"`
}

// LoadConfig create a new *Config and populate it with values from environment.
func LoadConfig(params LoadConfigParams) (Config, error) {
	cfg := Config{
		Prefix: "[X] ",
	}

	if params.Provider != nil {
		err := params.Provider.Get("xlog").Populate(&cfg)
		if err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}
