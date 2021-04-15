package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateVpcTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsResponse", string(data)}, " ")
}
