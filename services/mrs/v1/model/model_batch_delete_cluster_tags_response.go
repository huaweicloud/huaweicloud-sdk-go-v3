package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteClusterTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsResponse", string(data)}, " ")
}
