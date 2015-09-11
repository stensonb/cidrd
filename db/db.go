package db

import (
  "log"
  "errors"
  "github.com/jmoiron/sqlx"
  _ "github.com/mattn/go-sqlite3"
  "code.google.com/p/go-uuid/uuid"

  "github.com/stensonb/cidrd/config"
)

type DB struct {
  DB *sqlx.DB
}

func New(config *config.Config) (*DB, error) {
  // TODO: do something with config

  thedb := new(DB)
  var err error
  thedb.DB, err = sqlx.Open("sqlite3", ":memory:")
  if err != nil {
    return &DB{}, errors.New("couldn't open DB connection")
  }

  setupTestingDB(thedb)

  log.Printf("created memory db.")

  return thedb, nil
}

func setupTestingDB(db *DB) {
  // schema stuff
  schema := `CREATE TABLE class (
    uuid text not null primary key,
    created datetime not null,
    modified datetime not null,
    name text not null);`

  // execute a query on the server
  db.DB.MustExec(schema)

  theuuid := uuid.NewRandom().String()
  stuff := `INSERT INTO class VALUES ("` + theuuid + `", 1441922589, 1441922589, "bary")`
  db.DB.MustExec(stuff)
  stuff = `INSERT INTO class VALUES ("` + theuuid + `w", 1441923589, 1441923589, "emily")`
  db.DB.MustExec(stuff)
}
