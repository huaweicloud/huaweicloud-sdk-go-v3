package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupRequest Request Object
type ListPolicyGroupRequest struct {

	// 用于分页查询，返回策略组数量的限制。如果不指定，则返回所有符合条件的策略组。范围0~100。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，范围0~1000。
	Offset *int32 `json:"offset,omitempty"`

	// 根据策略组ID过滤结果。
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// 根据策略组名字过滤结果。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 根据优先级过滤结果。所带的值需要满足现有策略组已有最大优先级值。
	Priority *int32 `json:"priority,omitempty"`

	// 根据更新时间过滤结果。时间格式满足：yyyy-MM-dd HH:mm:ss。
	UpdateTime *string `json:"update_time,omitempty"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

	// 策略组名字精确查询。
	IsGroupNameAccurate *bool `json:"is_group_name_accurate,omitempty"`
}

func (o ListPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupRequest", string(data)}, " ")
}
