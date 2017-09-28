package room

import (
	"CrowdJam/models"
	"github.com/psyb0t/simplehttp"
)

// Updates room data
func UpdateRoom(r *simplehttp.Route) {
	room := models.NewRoom(r.Service.DB)
	room.SetId(r.Params.ByName("room_id"))

	err := room.Update(r.Input)
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	r.SuccessResponse()
}
