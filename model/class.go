package model

import (
  "time"
  //"errors"
  "code.google.com/p/go-uuid/uuid"
)

type Class struct {
  // boilerplate fields
  Uuid string         `json:"uuid"`
	Created time.Time		`json:"created"`
	Modified time.Time	`json:"modified"`

  Name string         `json:"name"`
}

// TODO: this goes away when pulling from sql db
var classes []*Class

func init() {
  // TODO: this goes away when pulling from sql db
  classes = []*Class{}
}

func (model *Model) Update(class *Class) error {
  // now validate what we're about to save
  err := validateUpdate(class)
  if err != nil {
    return err
  }

  err = model.save(class)
  if err != nil {
    return err
  }

  return nil
}

func validateUpdate(class *Class) error {
  return nil
}

func (class *Class) SetUUID(id string) {
  class.Uuid = id
}

func (model *Model) Create(class *Class) error {
  class.Uuid = generateID()

  // now validate what we're about to save
  err := validateCreate(class)
  if err != nil {
    return err
  }

  err = model.save(class)
  if err != nil {
    return err
  }

  return nil
}

// save this class into the database
// return non-nil error on failure
func (model *Model) save(c *Class) error {
  // created and modified can never be specified
  now := time.Now()
  c.Modified = now

  anewclass := new(Class)
  if c.Created == anewclass.Created {
    c.Created = now
  }

  // TODO: handle ID conflict here...generate another one, try again...
  // _, err :=
  _, err := model.db.DB.NamedExec(`INSERT INTO class VALUES (:uuid,:created,:modified,:name)`, c)
  if err != nil {
    return err
  }

  return nil
}

func generateID() string {
  return uuid.NewRandom().String()
}

func validateCreate(class *Class) error {
  return nil
}

func (model *Model) GetClassByUUID(id string) (*Class, error) {
  answer := Class{}
  err := model.db.DB.Get(&answer, "SELECT * FROM class where uuid=$1",id)
  if err != nil {
    return &Class{}, err
  }

  return &answer, nil
}

func (model *Model) GetAllClasses() (*[]Class, error) {
  lclasses := []Class{}

  err := model.db.DB.Select(&lclasses, "SELECT * FROM class")
  if err != nil {
    return &lclasses, err
  }

  return &lclasses, nil
}

func (model *Model) DeleteClassByUUID(id string) error {
  _, err := model.db.DB.NamedExec(`DELETE FROM class WHERE uuid=:uuid`, map[string]interface{}{"uuid": id,})
  if err != nil {
    return err
  }

  return nil
}
