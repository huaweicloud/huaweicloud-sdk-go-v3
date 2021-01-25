package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateTagResponse", string(data)}, " ")
}
