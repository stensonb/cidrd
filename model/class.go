package model

import (
	"time"
	//"errors"
)

type Class struct {
	// boilerplate fields
	Uuid     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`

	Name string `json:"name"`
}

// save this class into the database
// return non-nil error on failure
func (model *Model) StoreClass(c *Class) error {
	// updated modified time
	now := time.Now()
	c.Modified = now

	// if c.Created is nil/empty, then set it to now
	_emptyclass := new(Class)
	if c.Created == _emptyclass.Created {
		c.Created = now
	}

	// assign UUID if it doesn't exist
	if c.Uuid == _emptyclass.Uuid {
		c.Uuid = generateID()
	}

	// TODO: handle ID conflict here...generate another one, try again...

	_, err := model.GetClassByUUID(c.Uuid)
	if err != nil {
		// then the class doesn't exist, and we should insert it
		_, err = model.db.DB.NamedExec(`INSERT INTO class VALUES (:uuid,:created,:modified,:name)`, c)
		if err != nil {
			// TODO: if error is ID conflict, try generate another one and try again
			return err
		}
	} else {
		// the class exists, and we should update it
		_, err = model.db.DB.NamedExec(`UPDATE class SET uuid=:uuid,created=:created,modified=:modified,name=:name WHERE uuid=:uuid`, c)
		if err != nil {
			// TODO: if error is ID conflict, try generate another one and try again
			return err
		}
	}

	return nil
}

func (model *Model) GetClassByUUID(id string) (*Class, error) {
	answer := Class{}
	err := model.db.DB.Get(&answer, "SELECT * FROM class where uuid=$1", id)
	if err != nil {
		return &Class{}, err
	}

	return &answer, nil
}

func (model *Model) GetClassByName(name string) (*Class, error) {
	answer := Class{}
	err := model.db.DB.Get(&answer, "SELECT * FROM class where name=$1", name)
	if err != nil {
		return &Class{}, err
	}

	return &answer, nil
}

func (model *Model) GetAllClasses() (*[]Class, error) {
	lclasses := []Class{}

	err := model.db.DB.Select(&lclasses, "SELECT * FROM class")
	return &lclasses, err
}

func (model *Model) DeleteClassByUUID(id string) error {
	_, err := model.db.DB.NamedExec(`DELETE FROM class WHERE uuid=:uuid`, map[string]interface{}{"uuid": id})
	if err != nil {
		return err
	}

	return nil
}
