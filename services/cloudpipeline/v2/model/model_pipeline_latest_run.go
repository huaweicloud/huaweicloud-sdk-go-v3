package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineLatestRun 流水线及其最近一次运行信息
type PipelineLatestRun struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 流水线运行实例ID
	PipelineRunId *string `json:"pipeline_run_id,omitempty"`

	// 执行人ID
	ExecutorId *string `json:"executor_id,omitempty"`

	// 执行人名称
	ExecutorName *string `json:"executor_name,omitempty"`

	// 阶段状态信息
	StageStatusList *[]PipelineLatestRunStageStatusList `json:"stage_status_list,omitempty"`

	// 流水线状态
	Status *string `json:"status,omitempty"`

	// 运行序号
	RunNumber *int32 `json:"run_number,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	BuildParams *PipelineLatestRunBuildParams `json:"build_params,omitempty"`

	ArtifactParams *PipelineLatestRunArtifactParams `json:"artifact_params,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o PipelineLatestRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineLatestRun struct{}"
	}

	return strings.Join([]string{"PipelineLatestRun", string(data)}, " ")
}
