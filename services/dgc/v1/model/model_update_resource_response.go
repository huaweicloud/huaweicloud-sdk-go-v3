package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateResourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceResponse", string(data)}, " ")
}
