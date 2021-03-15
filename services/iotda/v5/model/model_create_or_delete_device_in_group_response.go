package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateOrDeleteDeviceInGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrDeleteDeviceInGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOrDeleteDeviceInGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateOrDeleteDeviceInGroupResponse", string(data)}, " ")
}
