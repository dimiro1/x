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

	"github.com/dimiro1/x/xutils"
	"os"
	"strconv"
	"time"
)

// Config holds options to configure the database sql.DB pool of connections
type Config struct {
	// DriverNams can be one of sqlite, mysql or postgres
	DriverName string
	DSN        string

	MaxIdleConns    *int
	MaxOpenConns    *int
	ConnMaxLifetime *time.Duration
}

// LoadConfig load configuration from env vars and create a new config struct to hold the values.
//
// X_DB_DRIVER_NAME sets the driver name, it is required;
// X_DB_DSN sets the database connection dsn, it is required;
// X_DB_MAX_IDLE_CONNS sets the max idle connections, it is not required;
// X_DB_MAX_OPEN_CONNS sets the max open connections, it is not required;
// X_DB_MAX_CONN_LIFETIME sets the connection lifetime, ot is not required.
func LoadConfig() (*Config, error) {
	var (
		cfg        = &Config{}
		driverName = xutils.GetenvDefault("X_DB_DRIVER_NAME", "")
		dsn        = xutils.GetenvDefault("X_DB_DSN", "")
	)

	if driverName == "" {
		return nil, errors.New("X_DB_DRIVER_NAME env var is required")
	}

	if dsn == "" {
		return nil, errors.New("X_DB_DSN env var is required")
	}

	if maxIdleConns, ok := os.LookupEnv("X_DB_MAX_IDLE_CONNS"); ok {
		i, err := strconv.Atoi(maxIdleConns)
		if err != nil {
			return nil, err
		}

		cfg.MaxIdleConns = &i
	}

	if maxOpenConns, ok := os.LookupEnv("X_DB_MAX_OPEN_CONNS"); ok {
		i, err := strconv.Atoi(maxOpenConns)
		if err != nil {
			return nil, err
		}

		cfg.MaxOpenConns = &i
	}

	if connMaxLifetime, ok := os.LookupEnv("X_DB_MAX_CONN_LIFETIME"); ok {
		d, err := time.ParseDuration(connMaxLifetime)
		if err != nil {
			return nil, err
		}

		cfg.ConnMaxLifetime = &d
	}

	cfg.DriverName = driverName
	cfg.DSN = dsn
	return cfg, nil
}
