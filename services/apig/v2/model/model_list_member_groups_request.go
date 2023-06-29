package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMemberGroupsRequest Request Object
type ListMemberGroupsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 字典编码。  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道后端云服务组的名称
	MemberGroupName *string `json:"member_group_name,omitempty"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  当前支持member_group_name。
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListMemberGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMemberGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListMemberGroupsRequest", string(data)}, " ")
}
