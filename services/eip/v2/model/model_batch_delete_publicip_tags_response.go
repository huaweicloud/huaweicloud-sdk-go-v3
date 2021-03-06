package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeletePublicipTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePublicipTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsResponse", string(data)}, " ")
}
