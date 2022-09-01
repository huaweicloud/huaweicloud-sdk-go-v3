package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBackendInstancesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id" xml:"vpc_channel_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 云服务器的名称
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o ListBackendInstancesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackendInstancesV2Request struct{}"
	}

	return strings.Join([]string{"ListBackendInstancesV2Request", string(data)}, " ")
}
