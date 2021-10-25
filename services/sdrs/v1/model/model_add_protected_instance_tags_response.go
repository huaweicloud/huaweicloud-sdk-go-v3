package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddProtectedInstanceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddProtectedInstanceTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceTagsResponse", string(data)}, " ")
}
