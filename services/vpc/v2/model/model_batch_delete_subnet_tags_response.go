package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSubnetTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsResponse", string(data)}, " ")
}
