package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDeviceGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeviceGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupResponse", string(data)}, " ")
}
