package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSubnetResponse struct {
	Subnet         *UpdateSubnetResponseObject `json:"subnet,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateSubnetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponse", string(data)}, " ")
}
