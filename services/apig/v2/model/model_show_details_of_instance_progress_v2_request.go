package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfInstanceProgressV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o ShowDetailsOfInstanceProgressV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceProgressV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceProgressV2Request", string(data)}, " ")
}
