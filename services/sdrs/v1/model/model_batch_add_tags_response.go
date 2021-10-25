package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddTagsResponse", string(data)}, " ")
}
