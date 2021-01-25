package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteResetTracksTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResetTracksTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskResponse", string(data)}, " ")
}
