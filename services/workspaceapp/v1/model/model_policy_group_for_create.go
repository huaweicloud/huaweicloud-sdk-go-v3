package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupForCreate struct {

	// 策略组名称，名称需要满足如下规则 1. 由英文、数字或者下划线组成，不能有空格 2. 字符长度范围1-55
	PolicyGroupName string `json:"policy_group_name"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

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
