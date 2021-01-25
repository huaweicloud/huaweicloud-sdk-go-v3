package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePredefineTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"UpdatePredefineTagsResponse", string(data)}, " ")
}
