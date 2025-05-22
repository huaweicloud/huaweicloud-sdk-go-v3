package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildStageRecord 构建步骤
type BuildStageRecord struct {

	// 步骤编号
	Id *string `json:"id,omitempty"`

	// 步骤状态
	Status *string `json:"status,omitempty"`

	// 状态码
	StatusCode *int32 `json:"status_code,omitempty"`

	// 日志状态
	LogStatus *string `json:"log_status,omitempty"`

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

	// 构建时长
	Duration *int32 `json:"duration,omitempty"`

	// 子任务构建耗时
	BuildDuration *int32 `json:"build_duration,omitempty"`

	// 等待时间
	PendingDuration *int32 `json:"pending_duration,omitempty"`

	// 构建记录ID
	BuildRecordId *string `json:"build_record_id,omitempty"`

	// 八爪鱼任务ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 步骤名称
	ExecutionStageName *string `json:"execution_stage_name,omitempty"`

	// 步骤名称
	DisplayName *string `json:"display_name,omitempty"`

	// 节点ID
	NodeId *int32 `json:"node_id,omitempty"`

	// 序号
	Sequence *int32 `json:"sequence,omitempty"`
}

func (o BuildStageRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildStageRecord struct{}"
	}

	return strings.Join([]string{"BuildStageRecord", string(data)}, " ")
}
