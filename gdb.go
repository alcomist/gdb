package gdb

import (
	"database/sql"
)

type DB struct {
	*sql.DB
	driverName string
	unsafe     bool
	//Mapper     *reflectx.Mapper
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db, driverName: driverName}, err
}
