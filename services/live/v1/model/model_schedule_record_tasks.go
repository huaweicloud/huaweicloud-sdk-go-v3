package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleRecordTasks struct {

	// 推流域名
	Domain string `json:"domain"`

	// 应用名称
	App string `json:"app"`

	// 流名称
	Stream string `json:"stream"`

	// 录制任务开始时间，Unix时间戳。如果不填表示立即启动录制。EndTime - StartTime不能超过24小时。
	StartTime *int64 `json:"start_time,omitempty"`

	// 录制任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且小于当前时间+7天。
	EndTime int64 `json:"end_time"`

	// 录制模板ID，对应模板必须存在否则会启动失败。
	TemplateId string `json:"template_id"`

	// 任务停止时间，当前还在生效的任务此值为0，任务未被停止但是超过了end_time后此值为自动停止时间。
	StopTime *int64 `json:"stop_time,omitempty"`

	// 录制任务ID
	TaskId *string `json:"task_id,omitempty"`
}

func (o ScheduleRecordTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleRecordTasks struct{}"
	}

	return strings.Join([]string{"ScheduleRecordTasks", string(data)}, " ")
}
