package model

import (
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
