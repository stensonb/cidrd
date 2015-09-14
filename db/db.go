package db

import (
	_ "code.google.com/p/go-uuid/uuid"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"

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
    uuid string not null primary key,
    created datetime not null,
    modified datetime not null,
    name string not null);`

	// execute a query on the server
	db.DB.MustExec(schema)

	theuuid := "123423afsasdf" //uuid.NewRandom().String()
	stuff := `INSERT INTO class VALUES ("` + theuuid + `", 1441922589, 1441922589, "bary")`
	db.DB.MustExec(stuff)
	//  stuff = `INSERT INTO class VALUES ("` + theuuid + `w", 1441923589, 1441923589, "emily")`
	//  db.DB.MustExec(stuff)

	schema = `CREATE TABLE netblock (
    uuid text not null primary key,
    created datetime not null,
    modified datetime not null,
    class_uuid string not null,
    starting_ip string not null,
    ending_ip string not null,
    foreign key(class_uuid) references class(uuid) ON UPDATE RESTRICT ON DELETE RESTRICT);`

	// the FK constraint above doesn't appear to work for sqlite3...grrr....

	// execute a query on the server
	db.DB.MustExec(schema)

	anotheruuid := "asdfasfd"
	stuff = `INSERT INTO netblock VALUES ("` + anotheruuid + `", 1441923589, 1441923589, "` + theuuid + `", "10.10.10.10", "10.20.30.40")`
	db.DB.MustExec(stuff)
}
