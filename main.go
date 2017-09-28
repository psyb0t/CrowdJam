package main

import (
	item_handlers "CrowdJam/handlers/item"
	room_handlers "CrowdJam/handlers/room"
	"CrowdJam/middleware"

	"github.com/psyb0t/simplehttp"
)

func main() {
	service := simplehttp.New("127.0.0.1:8585", "db")
	service.SetServerName("CrowdJam 0.1")

	// Rooms route group
	service.AddRouteGroup("/rooms",
		simplehttp.NewRouteGroupItem(
			"", "POST", room_handlers.CreateRoom,
			middleware.CleanupInput,
		),

		simplehttp.NewRouteGroupItem(
			"/:room_id", "GET", room_handlers.GetRoom,
		),

		simplehttp.NewRouteGroupItem(
			"/:room_id", "PUT", room_handlers.UpdateRoom,
			middleware.CleanupInput,
		),

		simplehttp.NewRouteGroupItem(
			"/:room_id", "DELETE", room_handlers.DeleteRoom, nil,
		),
	)
	// End rooms route group

	// Room items route group
	service.AddRouteGroup("/rooms/:room_id/items",
		simplehttp.NewRouteGroupItem(
			"", "GET", room_handlers.GetRoomItems, nil,
		),

		simplehttp.NewRouteGroupItem(
			"", "POST", item_handlers.CreateItem,
			middleware.CleanupInput,
		),

		simplehttp.NewRouteGroupItem(
			"/:item_id", "GET", item_handlers.GetItem, nil,
		),

		simplehttp.NewRouteGroupItem(
			"/:item_id", "PUT", item_handlers.UpdateItem,
			middleware.CleanupInput,
		),

		// Add queue system in the db neaparat
		simplehttp.NewRouteGroupItem(
			"/:item_id/vote", "PUT", item_handlers.VoteItem, nil,
		),

		simplehttp.NewRouteGroupItem(
			"/:id", "DELETE", item_handlers.DeleteItem, nil,
		),
	)
	// End room items route group

	service.Start()
}
