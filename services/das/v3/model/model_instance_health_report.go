package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceHealthReport struct {

	// 报告ID
	TaskId string `json:"task_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 实例名称
	InstanceName string `json:"instance_name"`

	// 诊断起始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 诊断终止时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`
}

func (o InstanceHealthReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceHealthReport struct{}"
	}

	return strings.Join([]string{"InstanceHealthReport", string(data)}, " ")
}
