package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfApiGroupV2Request struct {
	InstanceId string `json:"instance_id"`

	GroupId string `json:"group_id"`
}

func (o ShowDetailsOfApiGroupV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiGroupV2Request", string(data)}, " ")
}
