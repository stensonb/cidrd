package models

import (
  "time"
  "errors"
  "code.google.com/p/go-uuid/uuid"
//j  "github.com/jmoiron/sqlx"
//  "github.com/mattn/go-sqlite3"
)

type Class struct {
  // boilerplate fields
  UUID uuid.UUID			`json:"uuid"`
	Created time.Time		`json:"created"`
	Modified time.Time	`json:"modified"`

  Name string         `json:"name"`
}

type ClassList struct {
  Classes []*Class `json:"list"`
}

/*
type ClassDAO struct {
  DB *sqlx.DB
}
*/

var classes []*Class

func New() {
}

func Update(class *Class) error {
  class.Modified = time.Now()

  // now validate what we're about to save
  err := validateUpdate(class)
  if err != nil {
    return err
  }

  err = class.save()
  if err != nil {
    return err
  }

  return nil
}

func validateUpdate(class *Class) error {
  return nil
}

func (class *Class) SetUUID(id string) {
  class.UUID = uuid.Parse(id)
}

func Create(class *Class) error {
  class.UUID = generateID()

  // created and modified can never be specified
  now := time.Now()
  class.Created = now
  class.Modified = now

  // now validate what we're about to save
  err := validateCreate(class)
  if err != nil {
    return err
  }

  err = class.save()
  if err != nil {
    return err
  }

  return nil
}

// save this class into the database
// return non-nil error on failure
func (c *Class) save() error {
  // TODO: handle ID conflict here...generate another one, try again...
  classes = append(classes, c)
  return nil
}

func generateID() uuid.UUID {
  return uuid.NewRandom()
}

func validateCreate(class *Class) error {
  return nil
}

func GetClassByUUID(id string) (*Class, error) {
  // TODO: replace this with a sql query
  for i := range classes {
    if uuid.Equal(classes[i].UUID, uuid.Parse(id)) {
      return classes[i], nil
    }
  }

  // otherwise, we didn't find the class
  return &Class{}, errors.New("Not found.")
}

func GetAllClasses() (*ClassList, error) {
  ans := NewClassList(classes)
  return ans, nil
}

func NewClassList(classes []*Class) *ClassList {
  return &ClassList{classes}
}

func DeleteClassByUUID(id string) error {
  // TODO: replace this with a sql query
  for i := range classes {
    if uuid.Equal(classes[i].UUID, uuid.Parse(id)) {
      classes[i] = classes[len(classes)-1]
      classes = classes[:len(classes)-1]
      return nil
    }
  }

  // otherwise, we didn't find the
  return errors.New("Not found.")
}
