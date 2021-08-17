package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSpeedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSpeedResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateSpeedResponse", string(data)}, " ")
}
