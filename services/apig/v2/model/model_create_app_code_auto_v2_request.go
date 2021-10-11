package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAppCodeAutoV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
}

func (o CreateAppCodeAutoV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAppCodeAutoV2Request struct{}"
	}

	return strings.Join([]string{"CreateAppCodeAutoV2Request", string(data)}, " ")
}
