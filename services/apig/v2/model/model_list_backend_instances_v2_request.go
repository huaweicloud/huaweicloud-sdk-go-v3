package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBackendInstancesV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
	// 云服务器的名称

	Name *string `json:"name,omitempty"`
	// 每页显示的条数

	Limit *int32 `json:"limit,omitempty"`
	// 页码

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBackendInstancesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBackendInstancesV2Request struct{}"
	}

	return strings.Join([]string{"ListBackendInstancesV2Request", string(data)}, " ")
}
