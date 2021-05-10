package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateClusterTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsResponse", string(data)}, " ")
}
