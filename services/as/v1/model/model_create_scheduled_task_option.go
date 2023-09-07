package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateScheduledTaskOption struct {

	// 计划任务名称
	Name string `json:"name"`

	ScheduledPolicy *ScheduledTaskPolicy `json:"scheduled_policy"`

	InstanceNumber *IntegerRange `json:"instance_number"`
}

func (o CreateScheduledTaskOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduledTaskOption struct{}"
	}

	return strings.Join([]string{"CreateScheduledTaskOption", string(data)}, " ")
}
