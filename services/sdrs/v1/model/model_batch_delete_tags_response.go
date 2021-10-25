package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsResponse", string(data)}, " ")
}
