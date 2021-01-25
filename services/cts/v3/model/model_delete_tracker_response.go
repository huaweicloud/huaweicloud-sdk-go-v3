package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrackerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTrackerResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrackerResponse", string(data)}, " ")
}
