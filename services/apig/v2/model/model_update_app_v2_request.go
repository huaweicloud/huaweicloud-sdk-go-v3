package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// APP的编号

	AppId string `json:"app_id"`

	Body *AppReq `json:"body,omitempty"`
}

func (o UpdateAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAppV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAppV2Request", string(data)}, " ")
}
