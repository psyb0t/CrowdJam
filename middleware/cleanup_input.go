package middleware

import (
	"encoding/json"
	"github.com/psyb0t/simplehttp"
)

// Clean request input from unwanted params
func CleanupInput(r *simplehttp.Route) {
	m := make(map[string]interface{})
	err := json.Unmarshal(r.Input, &m)
	if err != nil {
		return
	}

	delete(m, "id")

	result_input, err := json.Marshal(m)
	if err != nil {
		return
	}

	r.Input = result_input
}
