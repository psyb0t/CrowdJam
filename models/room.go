package models

import (
	"CrowdJam/utils"
	"encoding/json"
	"fmt"

	"CrowdJam/collections"
	"CrowdJam/events"
	"github.com/psyb0t/simplehttp"
	"log"
)

type Room struct {
	ID          string            `json:"id""`
	Name        string            `json:"name""`
	Items       collections.Items `json:"items""`
	db          *simplehttp.DB
	dbKeyPrefix string
}

func NewRoom(db *simplehttp.DB) *Room {
	room := &Room{db: db, dbKeyPrefix: "room_"}
	room.InitEventListeners()

	return room
}

func (r *Room) InitEventListeners() {
	items := r.Items

	eventListener := events.NewEventListener()

	eventListener.On(
		func() bool {
			if len(items) != len(r.Items) {
				items = r.Items
				return true
			}

			return false
		},
	).Trigger("RoomItemsNumberChanged").HandledBy(
		func(ev *events.Event) {
			log.Println(ev.Name)
		},
	)

	eventListener = events.NewEventListener()

	eventListener.On(
		func() bool {
			for i, v := range items {
				if v != r.Items[i] {
					items = r.Items
					return true
				}
			}

			return false
		},
	).Trigger("RoomItemsValueChanged").HandledBy(
		func(ev *events.Event) {
			log.Println(ev.Name)
		},
	)

}

func (r *Room) SetId(id string) {
	r.ID = id
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
