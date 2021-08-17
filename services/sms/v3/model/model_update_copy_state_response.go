package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCopyStateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCopyStateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCopyStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCopyStateResponse", string(data)}, " ")
}
