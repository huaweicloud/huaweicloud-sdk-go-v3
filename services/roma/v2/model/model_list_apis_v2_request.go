package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApisV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// API编号
	Id *string `json:"id,omitempty" xml:"id"`

	// API名称
	Name *string `json:"name,omitempty" xml:"name"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 请求协议
	ReqProtocol *string `json:"req_protocol,omitempty" xml:"req_protocol"`

	// 请求方法
	ReqMethod *string `json:"req_method,omitempty" xml:"req_method"`

	// 请求路径
	ReqUri *string `json:"req_uri,omitempty" xml:"req_uri"`

	// 授权类型
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// API类型
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  目前仅支持name，req_uri，vpc_channel_name。
	PreciseSearch *string `json:"precise_search,omitempty" xml:"precise_search"`

	// 负载通道编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty" xml:"vpc_channel_id"`

	// 负载通道名称。
	VpcChannelName *string `json:"vpc_channel_name,omitempty" xml:"vpc_channel_name"`
}

func (o ListApisV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisV2Request struct{}"
	}

	return strings.Join([]string{"ListApisV2Request", string(data)}, " ")
}
