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

package xdb

import (
	"errors"

	"go.uber.org/config"
	"go.uber.org/fx"
	"time"
)

// Config holds options to configure the database sql.DB pool of connections
type Config struct {
	// DriverName can be one of sqlite, mysql or postgres
	DriverName string           `yaml:"driver"`
	DSN        string           `yaml:"dsn"`
	Connection ConfigConnection `yaml:"connection"`
}

type ConfigConnection struct {
	MaxIdle     int           `yaml:"max_idle"`
	MaxOpen     int           `yaml:"max_open"`
	MaxLifetime time.Duration `yaml:"max_lifetime"`
}

type LoadConfigParams struct {
	fx.In

	Provider config.Provider `optional:"true"`
}

func LoadConfig(params LoadConfigParams) (*Config, error) {
	cfg := &Config{
		DriverName: "",
		DSN:        "",
		Connection: ConfigConnection{
			MaxIdle:     0,
			MaxLifetime: 0,
			MaxOpen:     0,
		},
	}

	if params.Provider != nil {
		err := params.Provider.Get("xdb").Populate(&cfg)
		if err != nil {
			return nil, err
		}
	}

	if cfg.DriverName == "" {
		return nil, errors.New("xdb.driver is required")
	}

	if cfg.DSN == "" {
		return nil, errors.New("xdb.dsn is required")
	}

	return cfg, nil
}
