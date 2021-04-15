package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVpcTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpcTagsResponse", string(data)}, " ")
}
