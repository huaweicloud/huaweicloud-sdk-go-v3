package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupForUpdate struct {

	// 策略组名称。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

	// 策略来源。
	ScopeFlag *int32 `json:"scope_flag,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`

	Policies *Policies `json:"policies,omitempty"`
}

func (o PolicyGroupForUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupForUpdate struct{}"
	}

	return strings.Join([]string{"PolicyGroupForUpdate", string(data)}, " ")
}
