package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelingAuthorizationV2Request struct {
	InstanceId string `json:"instance_id"`
	AppAuthId  string `json:"app_auth_id"`
}

func (o CancelingAuthorizationV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelingAuthorizationV2Request struct{}"
	}

	return strings.Join([]string{"CancelingAuthorizationV2Request", string(data)}, " ")
}
