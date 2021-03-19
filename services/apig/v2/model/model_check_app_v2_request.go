package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckAppV2Request struct {
	InstanceId string `json:"instance_id"`

	AppId string `json:"app_id"`
}

func (o CheckAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckAppV2Request struct{}"
	}

	return strings.Join([]string{"CheckAppV2Request", string(data)}, " ")
}
