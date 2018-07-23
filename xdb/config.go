package xdb

import (
	"errors"

	"github.com/dimiro1/x/xutils"
	"os"
	"strconv"
	"time"
)

type Config struct {
	// DriverNams can be one of sqlite, mysql or postgres
	DriverName string
	DSN        string

	MaxIdleConns    *int
	MaxOpenConns    *int
	ConnMaxLifetime *time.Duration
}

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
