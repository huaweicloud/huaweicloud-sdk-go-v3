package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAuthorizingAppsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *AppAuthReq `json:"body,omitempty"`
}

func (o CreateAuthorizingAppsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAuthorizingAppsV2Request struct{}"
	}

	return strings.Join([]string{"CreateAuthorizingAppsV2Request", string(data)}, " ")
}
