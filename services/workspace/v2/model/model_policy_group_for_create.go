package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupForCreate struct {

	// 策略组id。
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// 策略组名称。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

	// 策略来源。
	ScopeFlag *int32 `json:"scope_flag,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`

	Policies *Policies `json:"policies,omitempty"`
}

func (o PolicyGroupForCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupForCreate struct{}"
	}

	return strings.Join([]string{"PolicyGroupForCreate", string(data)}, " ")
}
