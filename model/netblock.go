package model

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/stensonb/cidrd/model/netblockmath"
)

type Netblock struct {
	// boilerplate fields
	Uuid     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`

	Class_uuid  string `json:"class_uuid"`
	Starting_ip string `json:"starting_ip"`
	Ending_ip   string `json:"ending_ip"`
}

func (model *Model) validateNetblock(nb *Netblock) error {
	// ensure startingip and endingip are ip addresses
	startingip := net.ParseIP(nb.Starting_ip)
	if startingip == nil {
		return errors.New("startingip is not valid")
	}

	endingip := net.ParseIP(nb.Ending_ip)
	if endingip == nil {
		return errors.New("endingip is not valid")
	}

	// ensure the proposed netblock doesn't overlap
	// with any other existing netblock in the same
	// class.  for example, "rpcops" cannot have two
	// netblocks overlapping, but "rpcops" and "rpcfoo"
	// can have identical netblocks
	// get netblocks by class

	// TODO: Not threadsafe...multiple actors will mess this up
	proposedClass, err := model.GetClassByUUID(nb.Class_uuid)
	if err != nil {
		return err
	}

	nbids, err := model.GetNetblockIDsByClass(proposedClass.Name)
	if err != nil {
		return err
	}

	// go get each netblock in this class
	// and ensure the proposed netblock
	// doesn't overlap
	for _, i := range *nbids {
		// go get netblock
		_nb, err := model.GetNetblockByUUID(i)
		if err != nil {
			return err
		}

    // parse the ip strings into IP structs
		_nb_start := net.ParseIP(_nb.Starting_ip)
		_nb_end := net.ParseIP(_nb.Ending_ip)

    // if the proposed netblock overlaps, raise an error
		if netblockmath.NetblocksOverlap(&startingip, &endingip, &_nb_start, &_nb_end) {
			msg := fmt.Sprintf("netblock overlaps with %+s", _nb.Uuid)
			return errors.New(msg)
		}
	}

	return nil
}

// save this netblock into the database
// return non-nil error on failure
func (model *Model) StoreNetblock(n *Netblock) error {
	// validate netblock
	err := model.validateNetblock(n)
	if err != nil {
		return err
	}

	// updated modified time
	now := time.Now()
	n.Modified = now

	// if n.Created is nil/empty, then set it to now
	_emptynetblock := new(Netblock)
	if n.Created == _emptynetblock.Created {
		n.Created = now
	}

	// assign UUID if it doesn't exist
	if n.Uuid == _emptynetblock.Uuid {
		n.Uuid = generateID()
	}

	// TODO: handle ID conflict here...generate another one, try again...

	_, err = model.GetNetblockByUUID(n.Uuid)
	if err != nil {
		// then the netblock doesn't exist, and we should insert it
		_, err = model.db.DB.NamedExec(`INSERT INTO netblock VALUES (:uuid,:created,:modified,:class_uuid,:starting_ip,:ending_ip)`, n)
		if err != nil {
			// TODO: if error is ID conflict, try generate another one and try again
			return err
		}
	} else {
		// the netblock exists, and we should update it
		_, err = model.db.DB.NamedExec(`UPDATE netblock SET uuid=:uuid,created=:created,modified=:modified,class_uuid=:class_uuid,starting_ip=:starting_up,ending_ip=:ending_ip WHERE uuid=:uuid`, n)
		if err != nil {
			// TODO: if error is ID conflict, try generate another one and try again
			return err
		}
	}

	return nil
}

func (model *Model) GetNetblockByUUID(id string) (*Netblock, error) {
	answer := Netblock{}
	err := model.db.DB.Get(&answer, "SELECT * FROM netblock where uuid=$1", id)
	if err != nil {
		return &Netblock{}, err
	}

	return &answer, nil
}

func (model *Model) GetNetblockIDsByClass(class string) (*[]string, error) {
	netblockIDs := []string{}

	err := model.db.DB.Select(&netblockIDs, "select netblock.uuid from netblock join class on netblock.class_uuid=class.uuid where class.name=$1", class)

	return &netblockIDs, err
}

func (model *Model) GetAllNetblocks() (*[]Netblock, error) {
	netblocks := []Netblock{}

	err := model.db.DB.Select(&netblocks, "SELECT * FROM netblock")
	return &netblocks, err
}

func (model *Model) DeleteNetblockByUUID(id string) error {
	_, err := model.db.DB.NamedExec(`DELETE FROM netblock WHERE uuid=:uuid`, map[string]interface{}{"uuid": id})
	if err != nil {
		return err
	}

	return nil
}
