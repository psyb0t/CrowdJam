package models

import (
	"encoding/json"
)

type Item struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Votes  int64       `json:"votes"`
	Data   interface{} `json:"data"`
	roomID string
}

func NewItem(roomID string) *Item {
	return &Item{roomID: roomID}
}

func (i *Item) SetID(id string) {
	i.ID = id
}

func (i *Item) Create(jsonData []byte) error {
	err := json.Unmarshal(jsonData, i)
	if err != nil {
		return err
	}

	return nil
}

func (i *Item) Update(jsonData []byte) error {
	return nil
}

func (i *Item) Save() error {
	return nil
}

func (i *Item) Sync() error {
	return nil
}

func (i *Item) Delete() error {
	return nil
}
