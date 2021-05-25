package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// APP的编号

	AppId string `json:"app_id"`
}

func (o DeleteAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAppV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAppV2Request", string(data)}, " ")
}
