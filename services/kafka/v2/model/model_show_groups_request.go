package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowGroupsRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`
}

func (o ShowGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGroupsRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupsRequest", string(data)}, " ")
}
