package xdb

import (
	healthDB "github.com/dimiro1/health/db"
	"github.com/dimiro1/x/xhealth"
)

// TODO: implement a clever way to decide which checker have to be used
func DBHealthCheck(cfg *Config, db DB) xhealth.CheckMapping {
	var checker healthDB.Checker

	switch cfg.DriverName {
	case "sqlite3":
		checker = healthDB.NewSqlite3Checker(db.DB)
	case "postgres":
		checker = healthDB.NewPostgreSQLChecker(db.DB)
	case "mysql":
		checker = healthDB.NewMySQLChecker(db.DB)
	}

	return xhealth.CheckMapping{
		Checker: &xhealth.Checker{
			Name:    cfg.DriverName,
			Checker: checker,
		},
	}
}
