package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVpcChannelsV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// VPC通道的编号

	Id *string `json:"id,omitempty"`
	// VPC通道的名称

	Name *string `json:"name,omitempty"`
	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持

	DictCode *string `json:"dict_code,omitempty"`
	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  目前支持name，member_group_name。

	PreciseSearch *string `json:"precise_search,omitempty"`
	// 后端服务地址。默认精确查询，不支持模糊查询。

	MemberHost *string `json:"member_host,omitempty"`
	// 后端服务器端口

	MemberPort *int32 `json:"member_port,omitempty"`
	// 后端服务器组名称

	MemberGroupName *string `json:"member_group_name,omitempty"`
	// 后端服务器组编号

	MemberGroupId *string `json:"member_group_id,omitempty"`
}

func (o ListVpcChannelsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcChannelsV2Request struct{}"
	}

	return strings.Join([]string{"ListVpcChannelsV2Request", string(data)}, " ")
}
