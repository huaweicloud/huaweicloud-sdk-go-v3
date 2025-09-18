package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubTaskInfo 子任务进度信息
type SubTaskInfo struct {

	// 子任务名
	SubTaskName *string `json:"sub_task_name,omitempty"`

	// 子任务进度
	Percent *string `json:"percent,omitempty"`

	// 子任务执行时间
	ExecutedTime *string `json:"executed_time,omitempty"`

	// 子任务状态
	Status *string `json:"status,omitempty"`

	// 子任务剩余预估时长
	RemainingTime *string `json:"remaining_time,omitempty"`

	// 是否展示已恢复库表
	ShowDetail *bool `json:"show_detail,omitempty"`
}

func (o SubTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskInfo struct{}"
	}

	return strings.Join([]string{"SubTaskInfo", string(data)}, " ")
}
