package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsResponse", string(data)}, " ")
}
