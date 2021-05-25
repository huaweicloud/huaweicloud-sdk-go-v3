package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// APP的编号

	AppId string `json:"app_id"`
}

func (o ShowDetailsOfAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppV2Request", string(data)}, " ")
}
