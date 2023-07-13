package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScheduledTaskOption struct {

	// 计划任务名称
	Name *string `json:"name,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	InstanceNumber *IntegerRange `json:"instance_number,omitempty"`
}

func (o UpdateScheduledTaskOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTaskOption struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTaskOption", string(data)}, " ")
}
