package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupResponse", string(data)}, " ")
}
