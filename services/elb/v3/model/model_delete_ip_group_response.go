package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteIpGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIpGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteIpGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupResponse", string(data)}, " ")
}
