package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroupForUpdate 需要修改的协议策略组。
type PolicyGroupForUpdate struct {

	// 策略组名称，长度不能超过255个字符。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 策略组描述，长度不能超过255个字符。
	Description *string `json:"description,omitempty"`

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
