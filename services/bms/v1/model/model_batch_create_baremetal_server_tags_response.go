package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateBaremetalServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateBaremetalServerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateBaremetalServerTagsResponse", string(data)}, " ")
}
