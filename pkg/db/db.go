package db

import (
	// "github.com/xSaCh/dweep/pkg/mocks"
	// "github.com/xSaCh/dweep/pkg/models"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	connection *sqlx.DB
}

func NewDb() (*Db, error) {
	gdb, err := sqlx.Open("sqlite3", "test.db")

	if err != nil {
		// log.Fatalln("Error while opening db ", err)
		return nil, err
	}

	newDb := Db{connection: gdb}
	return &newDb, nil
}

func (*Db) Init() {
	// TODO: Create tables
	// Db.MustExec(scheme)
}
