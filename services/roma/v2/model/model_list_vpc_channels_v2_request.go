package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVpcChannelsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// VPC通道的编号
	Id *string `json:"id,omitempty" xml:"id"`

	// VPC通道的名称
	Name *string `json:"name,omitempty" xml:"name"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty" xml:"dict_code"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  目前支持name，member_group_name。
	PreciseSearch *string `json:"precise_search,omitempty" xml:"precise_search"`

	// 后端服务地址。默认精确查询，不支持模糊查询。
	MemberHost *string `json:"member_host,omitempty" xml:"member_host"`

	// 后端服务器端口
	MemberPort *int32 `json:"member_port,omitempty" xml:"member_port"`

	// 后端服务器组名称
	MemberGroupName *string `json:"member_group_name,omitempty" xml:"member_group_name"`

	// 后端服务器组编号
	MemberGroupId *string `json:"member_group_id,omitempty" xml:"member_group_id"`
}

func (o ListVpcChannelsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcChannelsV2Request struct{}"
	}

	return strings.Join([]string{"ListVpcChannelsV2Request", string(data)}, " ")
}
