package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCommandResultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCommandResultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCommandResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResultResponse", string(data)}, " ")
}
