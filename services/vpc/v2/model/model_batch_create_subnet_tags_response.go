package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSubnetTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsResponse", string(data)}, " ")
}
