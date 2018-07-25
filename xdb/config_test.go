package xdb

import (
	"os"
	"testing"
	"time"
)

func TestLoadConfig_Missing_Driver(t *testing.T) {
	os.Clearenv()

	c, err := LoadConfig()
	if err == nil {
		t.Errorf("err == %v, expected %v", err, nil)
	}

	if c != nil {
		t.Errorf("c == %v, expected !%v", c, nil)
	}
}

func TestLoadConfig_Missing_DSN(t *testing.T) {
	os.Clearenv()
	os.Setenv("X_DB_DRIVER_NAME", "sqlite3")

	c, err := LoadConfig()
	if err == nil {
		t.Errorf("err == %v, expected %v", err, nil)
	}

	if c != nil {
		t.Errorf("c == %v, expected !%v", c, nil)
	}
}

func TestLoadConfig(t *testing.T) {
	os.Clearenv()
	os.Setenv("X_DB_DRIVER_NAME", "sqlite3")
	os.Setenv("X_DB_DSN", ":memory:")

	os.Setenv("X_DB_MAX_IDLE_CONNS", "1")
	os.Setenv("X_DB_MAX_OPEN_CONNS", "2")
	os.Setenv("X_DB_MAX_CONN_LIFETIME", "10s")

	c, err := LoadConfig()
	if err != nil {
		t.Errorf("err == %v, expected !%v", err, nil)
	}

	if c.DriverName != "sqlite3" {
		t.Errorf("DriverName == %v, expected %v", c.DriverName, "sqlite3")
	}

	if c.DSN != ":memory:" {
		t.Errorf("DSN == %v, expected %v", c.DSN, ":memory:")
	}

	if *c.MaxIdleConns != 1 {
		t.Errorf("MaxIdleConns == %v, expected %v", *c.MaxIdleConns, 1)
	}

	if *c.MaxOpenConns != 2 {
		t.Errorf("MaxOpenConns == %v, expected %v", *c.MaxOpenConns, 2)
	}

	if *c.ConnMaxLifetime != 10*time.Second {
		t.Errorf("ConnMaxLifetime == %v, expected %v", *c.ConnMaxLifetime, 10*time.Second)
	}
}
