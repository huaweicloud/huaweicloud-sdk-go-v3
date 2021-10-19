package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddOrRemoveResourceInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrRemoveResourceInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceResponse", string(data)}, " ")
}
