package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroup 策略组
type PolicyGroup struct {

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

	// 策略来源。
	ScopeFlag *int32 `json:"scope_flag,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`

	Policies *Policies `json:"policies,omitempty"`
}

func (o PolicyGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroup struct{}"
	}

	return strings.Join([]string{"PolicyGroup", string(data)}, " ")
}
