package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateOrUpdateTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateOrUpdateTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOrUpdateTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateTagsResponse", string(data)}, " ")
}
