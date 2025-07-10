package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupForList struct {

	// 策略组ID。
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// 策略组名称。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 更新日期。
	UpdateTime *string `json:"update_time,omitempty"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

	Policies *Policies `json:"policies,omitempty"`

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`
}

func (o PolicyGroupForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupForList struct{}"
	}

	return strings.Join([]string{"PolicyGroupForList", string(data)}, " ")
}
