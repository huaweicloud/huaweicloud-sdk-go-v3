package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MonitorTaskInfo 任务监控连接信息。
type MonitorTaskInfo struct {

	// Flink作业任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务实时监控指标。
	TaskProps *interface{} `json:"task_props,omitempty"`
}

func (o MonitorTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorTaskInfo struct{}"
	}

	return strings.Join([]string{"MonitorTaskInfo", string(data)}, " ")
}
