package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfGatewayResponseV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 响应编号

	ResponseId string `json:"response_id"`
}

func (o ShowDetailsOfGatewayResponseV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfGatewayResponseV2Request", string(data)}, " ")
}
