package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduledTaskDetail struct {

	// 计划任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 伸缩组ID
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	InstanceNumber *IntegerRange `json:"instance_number,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ScheduledTaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskDetail struct{}"
	}

	return strings.Join([]string{"ScheduledTaskDetail", string(data)}, " ")
}
