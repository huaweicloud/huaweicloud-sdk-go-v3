package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResettingAppSecretV2Request struct {
	InstanceId string `json:"instance_id"`

	AppId string `json:"app_id"`

	Body *AppSecretReq `json:"body,omitempty"`
}

func (o ResettingAppSecretV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResettingAppSecretV2Request struct{}"
	}

	return strings.Join([]string{"ResettingAppSecretV2Request", string(data)}, " ")
}
