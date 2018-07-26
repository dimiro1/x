package xdb

import (
	"database/sql"
	"github.com/dimiro1/health/db"
	"github.com/dimiro1/x/xlog"
	"testing"
)

func TestDBHealthCheck_MySQL(t *testing.T) {
	var (
		cfg = &Config{
			DriverName: "mysql",
		}
		sqlDB = DB{
			DB: &sql.DB{},
		}
	)

	mapping := DBHealthCheck(cfg, sqlDB, xlog.OptionalLogger{})

	if mapping.Checker.Name != "mysql" {
		t.Errorf("mapping.Checker.Name == %v, expected %v", mapping.Checker.Name, "mysql")
	}
}

func TestDBHealthCheck_Postgres(t *testing.T) {
	var (
		cfg = &Config{
			DriverName: "postgres",
		}
		sqlDB = DB{
			DB: &sql.DB{},
		}
	)

	mapping := DBHealthCheck(cfg, sqlDB, xlog.OptionalLogger{})

	if mapping.Checker.Name != "postgres" {
		t.Errorf("mapping.Checker.Name == %v, expected %v", mapping.Checker.Name, "postgres")
	}
}

func TestDBHealthCheck_Sqlite3(t *testing.T) {
	var (
		cfg = &Config{
			DriverName: "sqlite3",
		}
		sqlDB = DB{
			DB: &sql.DB{},
		}
	)

	mapping := DBHealthCheck(cfg, sqlDB, xlog.OptionalLogger{})

	if mapping.Checker == nil {
		t.Errorf("mapping.Checker == %v, expected not %v", mapping.Checker, nil)
	}

	if mapping.Checker.Name != "sqlite3" {
		t.Errorf("mapping.Checker.Name == %v, expected %v", mapping.Checker.Name, "sqlite3")
	}

	checker, ok := mapping.Checker.Checker.(db.Checker)
	if !ok {
		t.Errorf("mapping.Checker.Checker has to be an instance of db.Checker")
	}

	// Not going to check the library
	if checker.VersionSQL == "" {
		t.Errorf("checker.VersionSQL == %v, expected not empty", checker.VersionSQL)
	}

	if checker.CheckSQL == "" {
		t.Errorf("checker.CheckSQL == %v, expected not empty", checker.CheckSQL)
	}
}

func TestDBHealthCheck_UnknownDriver(t *testing.T) {
	var (
		cfg = &Config{
			DriverName: "fake_driver",
		}
		sqlDB = DB{
			DB: &sql.DB{},
		}
	)

	mapping := DBHealthCheck(cfg, sqlDB, xlog.OptionalLogger{})

	if mapping.Checker == nil {
		t.Errorf("mapping.Checker == %v, expected not %v", mapping.Checker, nil)
	}

	if mapping.Checker.Name != "fake_driver" {
		t.Errorf("mapping.Checker.Name == %v, expected %v", mapping.Checker.Name, "fake_driver")
	}

	checker, ok := mapping.Checker.Checker.(db.Checker)
	if !ok {
		t.Errorf("mapping.Checker.Checker has to be an instance of db.Checker")
	}

	if checker.VersionSQL != "" {
		t.Errorf("checker.VersionSQL == %v, expected empty", checker.VersionSQL)
	}

	if checker.CheckSQL == "" {
		t.Errorf("checker.CheckSQL == %v, expected not empty", checker.CheckSQL)
	}
}
