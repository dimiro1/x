package xdb

import (
	"database/sql"
	"database/sql/driver"
	"testing"
)

type Dummy struct {
	driver.Driver
	OpenFunc func(name string) (driver.Conn, error)
}

func (d *Dummy) Open(name string) (driver.Conn, error) {
	return d.OpenFunc(name)
}

type DummyConn struct {
	driver.Conn
}

func TestNewDB(t *testing.T) {
	cfg := &Config{
		DriverName: "dummy",
		DSN:        ":memory:",
	}

	sql.Register("dummy", &Dummy{
		OpenFunc: func(name string) (driver.Conn, error) {
			return &DummyConn{}, nil
		},
	})

	db, err := NewDB(cfg)
	if err != nil {
		t.Errorf("err == %v, expected !%v", err, nil)
	}

	if db == nil {
		t.Errorf("db == %v, expected !%v", nil, nil)
	}
}
