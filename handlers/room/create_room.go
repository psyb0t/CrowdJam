package room

import (
	"CrowdJam/models"

	"github.com/psyb0t/simplehttp"
)

// Creates a new room from the input data
func CreateRoom(r *simplehttp.Route) {
	room := models.NewRoom(r.Service.DB)
	err := room.Create(r.Input)
	if err != nil {
		r.ErrorResponse(err.Error())
		return
	}

	r.SuccessObjectResponse(room)
}
