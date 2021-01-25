package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConnectionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionResponse", string(data)}, " ")
}
