package xdb

import (
	"database/sql"
	"go.uber.org/fx"
)

type DB struct {
	fx.In

	DB *sql.DB `name:"x_sql_db"`
}

type DBQualifier struct {
	fx.Out

	DB *sql.DB `name:"x_sql_db"`
}

func NewDB(cfg *Config) (DBQualifier, error) {
	db, err := sql.Open(cfg.DriverName, cfg.DSN)
	if err != nil {
		return DBQualifier{}, err
	}

	if cfg.MaxIdleConns != nil {
		db.SetMaxIdleConns(*cfg.MaxIdleConns)
	}

	if cfg.MaxOpenConns != nil {
		db.SetMaxOpenConns(*cfg.MaxOpenConns)
	}

	if cfg.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(*cfg.ConnMaxLifetime)
	}

	if err := db.Ping(); err != nil {
		return DBQualifier{}, err
	}

	return DBQualifier{
		DB: db,
	}, nil
}
