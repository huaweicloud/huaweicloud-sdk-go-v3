package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectVpcChannelsV2Request Request Object
type ListProjectVpcChannelsV2Request struct {

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// VPC通道的编号
	Id *string `json:"id,omitempty"`

	// VPC通道的名称
	Name *string `json:"name,omitempty"`

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

	// 是否返回后端实例列表
	MembersReturn *bool `json:"members_return,omitempty"`
}

func (o ListProjectVpcChannelsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectVpcChannelsV2Request struct{}"
	}

	return strings.Join([]string{"ListProjectVpcChannelsV2Request", string(data)}, " ")
}
