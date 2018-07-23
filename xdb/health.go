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
	healthDB "github.com/dimiro1/health/db"
	"github.com/dimiro1/x/xhealth"
	"github.com/dimiro1/x/xlog"
)

// DBHealthCheck register a health check for the database.
// offically it suports only mysql, postgres and mysql. however it try to register the unknown anyways.
func DBHealthCheck(cfg *Config, db DB, logger xlog.OptionalLogger) xhealth.CheckMapping {
	var checker healthDB.Checker

	switch cfg.DriverName {
	case "sqlite3":
		checker = healthDB.NewSqlite3Checker(db.DB)
	case "postgres":
		checker = healthDB.NewPostgreSQLChecker(db.DB)
	case "mysql":
		checker = healthDB.NewMySQLChecker(db.DB)
	default:
		if xlog.IsProvided(logger) {
			logger.Logger.Printf(
				"registering an database that we dont know how to extract the database version")
		}

		checker = healthDB.NewChecker(
			"SELECT 1",
			`SELECT "DONT KNOW HOW TO GET THE DATABASE VERSION"`,
			db.DB,
		)
	}

	return xhealth.CheckMapping{
		Checker: &xhealth.Checker{
			Name:    cfg.DriverName,
			Checker: checker,
		},
	}
}
