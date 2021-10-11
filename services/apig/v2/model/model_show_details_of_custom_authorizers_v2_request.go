package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfCustomAuthorizersV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 自定义认证的ID

	AuthorizerId string `json:"authorizer_id"`
}

func (o ShowDetailsOfCustomAuthorizersV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfCustomAuthorizersV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfCustomAuthorizersV2Request", string(data)}, " ")
}
