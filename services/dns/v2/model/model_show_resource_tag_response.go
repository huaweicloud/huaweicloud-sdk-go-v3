package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowResourceTagResponse struct {
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceTagResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagResponse", string(data)}, " ")
}
