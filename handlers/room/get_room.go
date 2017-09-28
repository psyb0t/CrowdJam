package room

import (
	"CrowdJam/models"
	"github.com/psyb0t/simplehttp"
)

// Fetches the room from the db by id
func GetRoom(r *simplehttp.Route) {
	room := models.NewRoom(r.Service.DB)
	room.SetId(r.Params.ByName("room_id"))

	err := room.Sync()
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	r.SuccessObjectResponse(room)
}
