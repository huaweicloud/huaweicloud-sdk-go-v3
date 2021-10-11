package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAppCodeV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`

	Body *AppCodeCreate `json:"body,omitempty"`
}

func (o CreateAppCodeV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAppCodeV2Request struct{}"
	}

	return strings.Join([]string{"CreateAppCodeV2Request", string(data)}, " ")
}
