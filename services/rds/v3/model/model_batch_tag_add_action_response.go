package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchTagAddActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagAddActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTagAddActionResponse struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionResponse", string(data)}, " ")
}
