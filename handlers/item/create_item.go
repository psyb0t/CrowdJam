package item

import (
	"CrowdJam/models"
	"github.com/psyb0t/simplehttp"
)

func CreateItem(r *simplehttp.Route) {
	room := models.NewRoom(r.Service.DB)
	room.SetId(r.Params.ByName("room_id"))

	err := room.Sync()
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	item := models.NewItem(room.ID)
	err = item.Create(r.Input)
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	room.Items.Add(item)

	r.SuccessObjectResponse(item)
}
