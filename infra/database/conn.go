package database

import "github.com/jmoiron/sqlx"

func GetDBConn(databaseUrl string) *sqlx.DB {
	sqlxDb := sqlx.MustConnect("mysql", databaseUrl)

	return sqlxDb
}
