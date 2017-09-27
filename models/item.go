package models

type Item struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Votes int64       `json:"votes"`
	Data  interface{} `json:"data"`
}

func NewItem() *Item {
	return &Item{}
}

func (i *Item) SetRoomId(roomID string) {
	i.ID = roomID
}
