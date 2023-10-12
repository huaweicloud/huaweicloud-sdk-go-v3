package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupRequest Request Object
type ListPolicyGroupRequest struct {

	// 查询的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]
	Limit *int32 `json:"limit,omitempty"`

	// 根据策略组名字过滤结果。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 根据策略组类型名字过滤结果。
	PolicyGroupType *int32 `json:"policy_group_type,omitempty"`
}

func (o ListPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupRequest", string(data)}, " ")
}
