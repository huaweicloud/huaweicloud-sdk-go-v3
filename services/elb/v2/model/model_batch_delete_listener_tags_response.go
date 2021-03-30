package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteListenerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsResponse", string(data)}, " ")
}
