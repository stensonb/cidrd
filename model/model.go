package model

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/stensonb/cidrd/db"
)

type Model struct {
	db *db.DB
}

func New(db *db.DB) *Model {
	return &Model{db: db}
}

func (m *Model) Close() {
	// TODO: cleanup db correctly here for shutdown
}

func generateID() string {
	return uuid.NewRandom().String()
}
