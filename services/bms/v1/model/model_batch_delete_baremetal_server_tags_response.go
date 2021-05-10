package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteBaremetalServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteBaremetalServerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsResponse", string(data)}, " ")
}
