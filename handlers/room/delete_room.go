package room

import (
	"CrowdJam/models"
	"github.com/psyb0t/simplehttp"
)

// Deletes a room
func DeleteRoom(r *simplehttp.Route) {
	room := models.NewRoom(r.Service.DB)
	room.SetId(r.Params.ByName("room_id"))

	err := room.Delete()
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	r.SuccessResponse()
}
