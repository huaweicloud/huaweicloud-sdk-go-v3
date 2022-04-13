package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBackendInstancesV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// 云服务器的名称

	Name *string `json:"name,omitempty"`
}

func (o ListBackendInstancesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackendInstancesV2Request struct{}"
	}

	return strings.Join([]string{"ListBackendInstancesV2Request", string(data)}, " ")
}
