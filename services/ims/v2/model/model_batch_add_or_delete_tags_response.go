package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddOrDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrDeleteTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsResponse", string(data)}, " ")
}
