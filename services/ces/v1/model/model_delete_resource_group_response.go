package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupResponse", string(data)}, " ")
}
