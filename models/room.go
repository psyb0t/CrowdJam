package models

import (
	"CrowdJam/utils"
	"encoding/json"
	"fmt"

	"github.com/psyb0t/simplehttp"
)

type Room struct {
	ID          string  `json:"id""`
	Name        string  `json:"name""`
	Items       []*Item `json:"items""`
	db          *simplehttp.DB
	dbKeyPrefix string
}

func NewRoom(db *simplehttp.DB) *Room {
	return &Room{db: db, dbKeyPrefix: "room_"}
}

func (r *Room) SetId(roomID string) {
	r.ID = roomID
}

func (r *Room) dbId() string {
	return fmt.Sprintf("%s%s", r.dbKeyPrefix, r.ID)
}

func (r *Room) Create(jsonData []byte) error {
	err := json.Unmarshal(jsonData, r)
	if err != nil {
		return err
	}

	idLength := 4
	loops := 0

	for {
		loops++
		if loops > 50 {
			idLength++
		}

		r.ID = utils.GenerateRandomNumeric(idLength)

		if !r.db.KeyExists(r.dbId()) {
			break
		}
	}

	err = r.Save()
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) Update(jsonData []byte) error {
	err := json.Unmarshal(jsonData, r)
	if err != nil {
		return err
	}

	err = r.Save()
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) Save() error {
	rSerialized, err := json.Marshal(r)
	if err != nil {
		return err
	}

	err = r.db.Set(r.dbId(), rSerialized)

	if err != nil {
		return err
	}

	return nil
}

func (r *Room) Sync() error {
	roomDBData, err := r.db.Get(r.dbId())
	if err != nil {
		return err
	}

	err = json.Unmarshal(roomDBData, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) Delete() error {
	return r.db.Delete(r.dbId())
}
