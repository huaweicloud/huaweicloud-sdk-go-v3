package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleTask 调度计划结构体
type ScheduleTask struct {

	// 设备id数组
	DeviceIds []string `json:"device_ids"`

	// 任务执行的动作，当前支持SetProperties
	Action string `json:"action"`

	// 对应action的参数
	Paras *interface{} `json:"paras"`
}

func (o ScheduleTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTask struct{}"
	}

	return strings.Join([]string{"ScheduleTask", string(data)}, " ")
}
