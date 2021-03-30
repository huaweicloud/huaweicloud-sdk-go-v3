package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateListenerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsResponse", string(data)}, " ")
}
