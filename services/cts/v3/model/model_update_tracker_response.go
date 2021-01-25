package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTrackerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTrackerResponse struct{}"
	}

	return strings.Join([]string{"UpdateTrackerResponse", string(data)}, " ")
}
