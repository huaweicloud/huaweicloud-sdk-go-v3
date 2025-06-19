package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BriefRecordItems struct {

	// 构建任务的ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 构建时长
	Duration *int32 `json:"duration,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 构建下发时间
	ScheduleTime *string `json:"schedule_time,omitempty"`

	// 构建排队耗时
	QueuedTime *string `json:"queued_time,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 子任务构建耗时
	BuildDuration *int32 `json:"build_duration,omitempty"`

	// 等待时间
	PendingDuration *int32 `json:"pending_duration,omitempty"`

	// 所属的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo *int32 `json:"build_no,omitempty"`

	// 代码分支
	Branch *string `json:"branch,omitempty"`

	// commitId
	Revision *string `json:"revision,omitempty"`

	// 触发者名称
	TriggerName *string `json:"trigger_name,omitempty"`

	// 构建编号，每日从1开始
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`
}

func (o BriefRecordItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BriefRecordItems struct{}"
	}

	return strings.Join([]string{"BriefRecordItems", string(data)}, " ")
}
