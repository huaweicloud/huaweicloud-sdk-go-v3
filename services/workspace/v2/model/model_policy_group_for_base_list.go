package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroupForBaseList struct {

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
}

func (o PolicyGroupForBaseList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupForBaseList struct{}"
	}

	return strings.Join([]string{"PolicyGroupForBaseList", string(data)}, " ")
}
