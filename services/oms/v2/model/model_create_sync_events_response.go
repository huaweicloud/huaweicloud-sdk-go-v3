package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSyncEventsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSyncEventsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSyncEventsResponse struct{}"
	}

	return strings.Join([]string{"CreateSyncEventsResponse", string(data)}, " ")
}
